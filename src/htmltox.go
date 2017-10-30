/*
Package htmltox defines the HTML conversion API server that interfaces with the
Chrome browser
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	chrome "github.com/mkenney/go-chrome"
	api "github.com/mkenney/go-rest-lite"
)

/*
HTMLToX defines the struct for the HTML conversion API service
*/
type HTMLToX struct {
	browser *chrome.Browser
	sockets map[string]*chrome.Socket
}

/*
NewAPIService returns a pointer to an HTMLToX struct
*/
func NewAPIService() (*HTMLToX, error) {
	var err error
	htmltox := new(HTMLToX)
	err = chrome.Launch(0, "", "", "")
	if nil != err {
		return nil, err
	}
	htmltox.browser, err = chrome.GetBrowser()
	if nil != err {
		return nil, err
	}
	htmltox.sockets = make(map[string]*chrome.Socket)
	return htmltox, nil
}

/*
Render takes an HTML source, either a string or a URL, and returns
a byte array of the resulting image

@param source An HTML string or URL
@param format An output format, one of 'jpg', 'png', 'pdf'
@param width The viewport width
@param height The viewport height
*/
func (htmltox *HTMLToX) Render(source, format string, width, height int) (result []byte, err error) {
	return
}

/*
Listen starts the HTTP service
*/
func (htmltox *HTMLToX) Listen(port int) (err error) {
	server := api.NewServer()
	defineRoutes(server)
	server.ListenAndServe(fmt.Sprintf(":%d", port))
	return
}

/*
Define all the route handlers for the service
*/
func defineRoutes(server *api.API) {

	// Usage
	server.AddHandler("/", func(request *http.Request, response *api.Response) {
		if "GET" == request.Method {
			content, err := ioutil.ReadFile("testdata/hello")
			if err != nil {
				response.StatusCode = 500
				response.StatusMessage = fmt.Sprintf("%v", err)
				response.Errors = append(response.Errors, err)
				log.Error(err)
			} else {
				response.Channel <- content
			}
		}
		response.Channel <- response.Done()
	})

	// Image conversion
	server.AddHandler("/image", func(request *http.Request, response *api.Response) {
		if "POST" == request.Method {
			response.Channel <- "POST:/image"
		}
		response.Channel <- response.Done()
	})

	// PDF conversion
	server.AddHandler("/pdf", func(request *http.Request, response *api.Response) {
		if "POST" == request.Method {
			response.Channel <- "POST:/pdf"
		}
		response.Channel <- response.Done()
	})
}
