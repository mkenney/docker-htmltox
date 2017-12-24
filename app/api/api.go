/*
Package api provides the REST api service
*/
package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetFormatter(&log.TextFormatter{})
}

/*
API contains HTTP and SQL helper functions and manages pointers to those resources
*/
type API struct {
	router *mux.Router
}

/*
New initializes and returns a pointer to an API struct
*/
func New() *API {
	return &API{router: mux.NewRouter()}
}

/*
Run starts the HTTP listener process
*/
func (api *API) Run(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), api.router))
}

/*
Handle is a wrapper to add logging to gorilla/mux managed routes
This should be used for adding routes to the API service.
*/
func (api *API) Handle(method, path string, handler func(http.ResponseWriter, *http.Request)) (r *mux.Route) {
	wrapper := func(response http.ResponseWriter, request *http.Request) {
		log.Infof("%s: %s", request.Method, request.RequestURI)
		handler(response, request)
	}
	return api.router.HandleFunc(path, wrapper)
}

/*
NotFoundHandler is a wrapper to add a route not found handler to the mux router
*/
func (api *API) NotFoundHandler(handler func(http.ResponseWriter, *http.Request)) {
	api.router.NotFoundHandler = http.HandlerFunc(handler)
}

/*
RespondWithErrorBody returns a properly formed error response
A JSON body is used for all error responses
*/
func (api *API) RespondWithErrorBody(
	request *http.Request,
	response http.ResponseWriter,
	code int,
	payload interface{},
	headers map[string]string) {

	log.Debugf("%s: %s - Sending error response body", request.Method, request.RequestURI)
	if 300 > code {
		log.Errorf("%s: %s - '%d' is not a valid error response code!", request.Method, request.RequestURI, code)
	}
	api.RespondWithJSONBody(request, response, code, payload, headers)
}

/*
RespondWithJSONBody returns a JSON encoded HTTP response body
*/
func (api *API) RespondWithJSONBody(
	request *http.Request,
	response http.ResponseWriter,
	code int,
	payload interface{},
	headers map[string]string) {

	log.Debugf("%s: %s - Sending JSON encoded response body", request.Method, request.RequestURI)
	response.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(payload)
	if nil != err {
		log.Errorf("%s: %s - %s", request.Method, request.RequestURI, err)
		code = 500
		body = []byte(`["An unknown error occurred"]`)
	}
	if 300 >= code {
		addCacheHeaders(response, code)
	}
	sendResponse(request, response, code, string(body), headers)
}

/*
RespondWithEncodedBody returns a base64 encoded HTTP response body
*/
func (api *API) RespondWithEncodedBody(
	request *http.Request,
	response http.ResponseWriter,
	code int,
	body string,
	headers map[string]string) {

	log.Debugf("%s: %s - Sending base64 encoded response body", request.Method, request.RequestURI)
	content := base64.StdEncoding.EncodeToString([]byte(body))
	sendResponse(request, response, code, string(content), headers)
}

/*
RespondWithRawBody returns an unmodified HTTP response body
*/
func (api *API) RespondWithRawBody(
	request *http.Request,
	response http.ResponseWriter,
	code int,
	body string,
	headers map[string]string) {

	log.Debugf("%s: %s - Sending raw response body", request.Method, request.RequestURI)
	sendResponse(request, response, code, body, headers)
}

func sendResponse(
	request *http.Request,
	response http.ResponseWriter,
	code int,
	body string,
	headers map[string]string,
) {

	if 300 > code {
		addCacheHeaders(response, code)
	}

	log.Debugf("%s: %s - Setting 'Access-Control-Allow-Origin' to '*'", request.Method, request.RequestURI)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	for k, v := range headers {
		response.Header().Set(k, v)
	}

	if responseHeaders, err := json.Marshal(response.Header()); nil == err {
		log.Debugf("%s: %s - Response headers: %s", request.Method, request.RequestURI, responseHeaders)
	}

	log.Infof("%s: %s - %d %s", request.Method, request.RequestURI, code, http.StatusText(code))
	if _, err := response.Write([]byte(body)); nil != err {
		log.Error(err)
		if strings.Contains(err.Error(), "Content-Length") {
			log.Errorf("Route handler may have exited before response was sent")
		}
	}
}

/*
addCacheHeaders adds the cache-policy headers to an HTTP response
*/
func addCacheHeaders(response http.ResponseWriter, code int) {
	log.Debugf("Adding cache control headers")
	response.Header().Set("Cache-Control", "private, no-cache, must-revalidate, max-age=0, proxy-revalidate, s-maxage=0")
	response.Header().Set("Pragma", "no-cache")
	response.Header().Set("Expires", "0")
	if code < 300 || code == 304 {
		response.Header().Set("Vary", "*")
	}
}
