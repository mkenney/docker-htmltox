/*
Package htmltox defines the HTML conversion API server that interfaces with the
Chrome browser
*/
package htmltox

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/mkenney/docker-htmltox/app/api"
	chrome "github.com/mkenney/go-chrome"
	"github.com/mkenney/go-chrome/cdtp/emulation"
	"github.com/mkenney/go-chrome/cdtp/page"
	"github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
HTMLToX defines the struct for the HTML conversion API service
*/
type HTMLToX struct {
	Browser chrome.Chromium
	Sockets map[string]socket.Socketer
	API     *api.API
}

/*
New returns a pointer to an HTMLToX struct
*/
func New() (*HTMLToX, error) {
	var err error

	htmltox := &HTMLToX{
		API: api.New(),
		Browser: chrome.New(&chrome.Flags{
			"addr":               []interface{}{"localhost"},
			"disable-extensions": nil,
			"disable-gpu":        nil,
			"headless":           nil,
			"hide-scrollbars":    nil,
			"no-first-run":       nil,
			"no-sandbox":         nil,
			"port":               []interface{}{9222},
			"remote-debugging-address": []interface{}{"0.0.0.0"},
			"remote-debugging-port":    []interface{}{9222},
		}, "", "", "", ""),
		Sockets: make(map[string]socket.Socketer),
	}

	err = htmltox.Browser.Launch()
	if nil != err {
		log.Error(err)
		return nil, err
	}

	htmltox.API.Handle("GET", "/", htmltox.Usage)
	htmltox.API.Handle("GET", "/test", htmltox.RenderURL)
	htmltox.API.Handle("GET", "/favicon.ico", func(response http.ResponseWriter, request *http.Request) {
		data, err := ioutil.ReadFile("/go/src/github.com/mkenney/docker-htmltox/app/assets/favicon.ico")
		if nil != err {
			log.Debugf(err.Error())
			return
		}
		headers := make(map[string]string)
		headers["Content-Type"] = "image/vnd.microsoft.icon"
		htmltox.API.RespondWithRawBody(
			request,
			response,
			200,
			string(data),
			headers,
		)
	})

	return htmltox, nil
}

/*
Usage returns usage information
*/
func (htmltox *HTMLToX) Usage(response http.ResponseWriter, request *http.Request) {
	headers := make(map[string]string)
	content, err := ioutil.ReadFile("/go/src/github.com/mkenney/docker-htmltox/app/usage.html")
	if err != nil {
		log.Error(err)
		htmltox.API.RespondWithErrorBody(
			request,
			response,
			500,
			fmt.Sprintf("%s", err),
			headers,
		)
	} else {
		htmltox.API.RespondWithRawBody(
			request,
			response,
			200,
			string(content),
			headers,
		)
	}
}

/*
RenderURL takes a URL as the HTML source and returns a byte array of the resulting image

@param source An HTML string or URL
@param format An output format, one of 'jpg', 'png', 'pdf'
@param width The viewport width
@param height The viewport height
*/
func (htmltox *HTMLToX) RenderURL(response http.ResponseWriter, request *http.Request) {

	//switch request.Method {
	//case "GET":
	//case "POST":
	//default:
	//	log.Errorf("Invalid request method '%s'", request.Method)
	//	htmltox.API.RespondWithErrorBody(
	//		request,
	//		response,
	//		400,
	//		fmt.Sprintf("%s is not a valid request method", request.Method),
	//		headers,
	//	)
	//	return
	//}

	queryParams, err := getParams(request)
	tmp, _ := json.Marshal(queryParams)
	log.Debugf("Query params: %s", string(tmp))
	if nil != err {
		htmltox.API.RespondWithErrorBody(
			request,
			response,
			400,
			fmt.Errorf("%s", "Failed to parse query params"),
			make(map[string]string),
		)
		return
	}

	tab, err := htmltox.Browser.NewTab(queryParams["url"][0])
	if nil != err {
		log.Error(err)
		htmltox.API.RespondWithErrorBody(
			request,
			response,
			500,
			err.Error(),
			make(map[string]string),
		)
		return
	}

	// Enable Page events
	enableResult := <-tab.Page().Enable()
	if nil != enableResult.CDTPError {
		log.Errorf("Page.Enable: %s", enableResult.CDTPError.Error())
		htmltox.API.RespondWithErrorBody(
			request,
			response,
			500,
			fmt.Sprintf("%s", enableResult.CDTPError),
			make(map[string]string),
		)
		return
	}

	emulationSizeResult := <-tab.Emulation().SetVisibleSize(&emulation.SetVisibleSizeParams{
		Width:  1440,
		Height: 1440,
	})
	if nil != emulationSizeResult.CDTPError {
		log.Error(emulationSizeResult.CDTPError)
	}

	// Set the viewport stuff
	emulationDeviceResult := <-tab.Emulation().SetDeviceMetricsOverride(&emulation.SetDeviceMetricsOverrideParams{
		Width:  1440,
		Height: 1440,
		ScreenOrientation: &emulation.ScreenOrientation{
			Type:  "portraitPrimary",
			Angle: 90,
		},
	})
	if nil != emulationDeviceResult.CDTPError {
		log.Errorf("Emulation.SetDeviceMetricsOverride: %s", emulationDeviceResult.CDTPError.Error())
	}

	screenshotCaptureStarted := false
	screenshotCaptured := false
	screenshotReturned := make(chan bool)
	renderScreenshot := func() string {
		screenshotCaptureStarted = true
		result := <-tab.Page().CaptureScreenshot(&page.CaptureScreenshotParams{
			Format: queryParams["format"][0],
		})
		if nil != result.CDTPError {
			log.Errorf("Page.CaptureScreenshot: %s", result.CDTPError.Error())
			htmltox.API.RespondWithErrorBody(
				request,
				response,
				500,
				fmt.Sprintf("%s", err),
				make(map[string]string),
			)
			return ""
		}
		log.Debugf("Screenshot rendered")
		return result.Data
	}

	returnScreenshot := func(data string) {
		bytes, err := base64.StdEncoding.DecodeString(data)
		if nil != err {
			htmltox.API.RespondWithErrorBody(
				request,
				response,
				500,
				fmt.Sprintf("%s", err),
				make(map[string]string),
			)
		}
		headers := make(map[string]string)
		headers["Content-Type"] = fmt.Sprintf("image/%s", queryParams["format"][0])
		htmltox.API.RespondWithRawBody(
			request,
			response,
			200,
			string(bytes),
			headers,
		)
		screenshotReturned <- true
	}

	loadEventHandler := socket.NewEventHandler("Page.loadEventFired", func(response *socket.Response) {
		if false == screenshotCaptureStarted {
			returnScreenshot(renderScreenshot())
			screenshotCaptured = true
		}
	})
	tab.AddEventHandler(loadEventHandler)

	// Don't wait too long
	if "" == queryParams["timeout"][0] {
		queryParams["timeout"][0] = "30"
	}
	timeout, err := strconv.Atoi(queryParams["timeout"][0])
	if nil != err {
		htmltox.API.RespondWithErrorBody(
			request,
			response,
			500,
			fmt.Sprintf("%s", err),
			make(map[string]string),
		)
	}

	// Force a render after max time
	maxUntil := time.Now().Add(time.Second * time.Duration(timeout))
	for {
		if true == screenshotCaptured {
			break
		} else if true == screenshotCaptureStarted {
			if true == <-screenshotReturned {
				break
			}
		} else if time.Now().Before(maxUntil) {
			time.Sleep(1)
		} else {
			screenshotCaptureStarted = true
			returnScreenshot(renderScreenshot())
			screenshotCaptured = true
		}
	}
}

func getParams(request *http.Request) (url.Values, error) {
	params, err := url.ParseQuery(request.URL.RawQuery)
	if nil != err {
		return nil, err
	}

	// format
	// Must be either "png" or "jpeg". Default "png"
	if _, ok := params["format"]; !ok || 0 == len(params["format"]) {
		params["format"] = make([]string, 1)
		params["format"][0] = "png"
	} else if "jpg" == params["format"][0] {
		params["format"][0] = "jpeg"
	} else if "png" != params["format"][0] && "jpeg" != params["format"][0] && "pdf" != params["format"][0] {
		return nil, fmt.Errorf("Invalid format '%s', must be either 'png', 'jpeg' or 'pdf'", params["format"])
	} else if len(params["format"]) > 1 {
		return nil, fmt.Errorf("Only one 'format' parameter is allowed")
	}

	// height
	// Must be an integer. Must have only 1 value
	if _, ok := params["height"]; !ok || 0 == len(params["height"]) {
		params["height"] = make([]string, 1)
		params["height"][0] = ""
	} else if _, err := strconv.Atoi(params["height"][0]); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Invalid height '%s'", params["height"])
	} else if len(params["height"]) > 1 {
		return nil, fmt.Errorf("Only one 'height' parameter is allowed")
	}

	// quality
	// Only applicable to the "jpeg" format. Must be an integer. Must have only 1 value
	if _, ok := params["quality"]; "jpeg" == params["format"][0] && !ok {
		params["quality"] = make([]string, 1)
		params["quality"][0] = "100"
	} else if "jpeg" != params["format"][0] && len(params["quality"]) > 0 {
		return nil, fmt.Errorf("The 'quality' param only applies to the 'jpeg' format")
	} else if len(params["quality"]) > 0 {
		if _, err := strconv.Atoi(params["quality"][0]); err != nil {
			log.Error(err)
			return nil, fmt.Errorf("Invalid quality '%s'", params["quality"])
		} else if len(params["quality"]) > 1 {
			return nil, fmt.Errorf("Only one 'quality' parameter is allowed")
		}
	}

	// scale
	// Must be an integer. Must have only 1 value
	if _, ok := params["scale"]; !ok || 0 == len(params["scale"]) {
		params["scale"] = make([]string, 1)
		params["scale"][0] = "1"
	} else if _, err := strconv.Atoi(params["scale"][0]); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Invalid scale '%s'", params["scale"])
	} else if len(params["scale"]) > 1 {
		return nil, fmt.Errorf("Only one 'scale' parameter is allowed")
	}

	// timeout
	// Must be an integer. Must have only 1 value
	if _, ok := params["timeout"]; !ok || 0 == len(params["timeout"]) {
		params["timeout"] = make([]string, 1)
		params["timeout"][0] = ""
	} else if _, err := strconv.Atoi(params["timeout"][0]); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Invalid timeout '%s'", params["timeout"])
	} else if len(params["timeout"]) > 1 {
		return nil, fmt.Errorf("Only one 'timeout' parameter is allowed")
	}

	// url
	// Must be a valid URL. Multiple values allowed unless "raw" is specified
	if _, ok := params["url"]; !ok || 0 == len(params["url"]) {
		params["url"] = make([]string, 1)
		params["url"][0] = ""
	} else {
		for k, urlParam := range params["url"] {
			if "/" != params["url"][k][len(params["url"][k])-1:] {
				params["url"][k] += "/"
			}
			if _, err := url.ParseRequestURI(urlParam); nil != err {
				return nil, fmt.Errorf("Invalid URL '%s'", urlParam)
			}
		}
	}

	// width
	// Must be an integer. Must have only 1 value
	if _, ok := params["width"]; !ok || 0 == len(params["width"]) {
		params["width"] = make([]string, 1)
		params["width"][0] = ""
	} else if _, err := strconv.Atoi(params["width"][0]); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Invalid width '%s'", params["width"])
	} else if len(params["width"]) > 1 {
		return nil, fmt.Errorf("Only one 'width' parameter is allowed")
	}

	// x-offset
	// Must be an integer. Must have only 1 value
	if _, ok := params["x-offset"]; !ok || 0 == len(params["x-offset"]) {
		params["x-offset"] = make([]string, 1)
		params["x-offset"][0] = ""
	} else if _, err := strconv.Atoi(params["x-offset"][0]); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Invalid x-offset '%s'", params["x-offset"])
	} else if len(params["x-offset"]) > 1 {
		return nil, fmt.Errorf("Only one 'x-offset' parameter is allowed")
	}

	// y-offset
	// Must be an integer. Must have only 1 value
	if _, ok := params["y-offset"]; !ok || 0 == len(params["y-offset"]) {
		params["y-offset"] = make([]string, 1)
		params["y-offset"][0] = ""
	} else if _, err := strconv.Atoi(params["y-offset"][0]); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Invalid y-offset '%s'", params["y-offset"])
	} else if len(params["y-offset"]) > 1 {
		return nil, fmt.Errorf("Only one 'y-offset' parameter is allowed")
	}

	return params, nil
}

//func getHandler(params url.Values, api *api.API, response http.ResponseWriter) (func(results []chrome.SocketScreenshotResult), error) {
//	var raw bool
//	if _, ok := params["raw"]; ok {
//		raw = true
//	}
//	if len(params["url"]) > 1 && raw {
//		return nil, fmt.Errorf("'raw' is an invalid parameter when rendering multiple images")
//	} else if raw {
//		return func(results []chrome.SocketScreenshotResult) {
//			api.RespondWithImage(response, http.StatusOK, results[0].Data, "jpeg")
//			log.Debug("Rendered screenshot sent")
//		}, nil
//	} else {
//		return func(results []chrome.SocketScreenshotResult) {
//			data := make([]string, 0)
//			for _, result := range results {
//				data = append(data, result.Data)
//			}
//			api.RespondWithJSON(response, http.StatusOK, data)
//		}, nil
//	}
//}

//func getHandlerTest(params url.Values, api *api.API, response http.ResponseWriter) (func(results []chrome.SocketResult), error) {
//	var raw bool
//	if _, ok := params["raw"]; ok {
//		raw = true
//	}
//	if len(params["url"]) > 1 && raw {
//		return nil, fmt.Errorf("'raw' is an invalid parameter when rendering multiple images")
//	} else if raw {
//		return func(results []chrome.SocketResult) {
//			api.RespondWithImage(response, http.StatusOK, results[0].Data, "jpeg")
//			log.Debug("Result sent")
//		}, nil
//	} else {
//		return func(results []chrome.SocketResult) {
//			data := make([]string, 0)
//			for _, result := range results {
//				data = append(data, result.Data)
//			}
//			api.RespondWithJSON(response, http.StatusOK, data)
//		}, nil
//	}
//}

//func (htmltox *HTMLToX) test(response http.ResponseWriter, request *http.Request) {
//	var err error
//
//	switch request.Method {
//	case "GET":
//	case "POST":
//	default:
//		log.Errorf("Invalid request method '%s'", request.Method)
//		htmltox.API.RespondWithError(response, 400, fmt.Sprintf("%s is not a valid request method", request.Method))
//		return
//	}
//
//	params, err := getParams(request)
//	tmp, _ := json.Marshal(params)
//	log.Debugf("Query params: %s", string(tmp))
//	if nil != err {
//		log.Errorf("Failed to parse query params: %s", err)
//		htmltox.API.RespondWithError(response, 400, fmt.Sprintf("%s", err))
//		return
//	}
//
//	handler, err := getHandlerTest(params, htmltox.API, response)
//	if nil != err {
//		log.Errorf("Failed to generate response handler: %s", err)
//		htmltox.API.RespondWithError(response, 400, fmt.Sprintf("%s", err))
//		return
//	}
//	//chrome.RenderScreenshots(params, handler)
//	chrome.RenderScreenshotsTest(params, handler)
//
//	return
//}
