/*
Package api provides the REST api service
*/
package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
}

/*
The App struct contains HTTP and SQL helper functions and manages pointers to those resources
*/
type API struct {
	router *mux.Router
}

/*
New initializes and returns a pointer to an API struct
*/
func New() *API {
	a := new(API)
	a.router = mux.NewRouter()
	return a
}

/*
Run starts the HTTP listener process
*/
func (a *API) Run(port int) {
	log.Fatal(http.ListenAndServe(":80", a.router))
}

/*
Handle is a wrapper to add logging to gorilla/mux managed routes

This should be used for adding routes to the API service.
*/
func (a *API) Handle(method, path string, handler func(http.ResponseWriter, *http.Request)) (r *mux.Route) {
	wrapper := func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s", r.Method, r.RequestURI)
		handler(rw, r)
	}
	return a.router.HandleFunc(path, wrapper)
}

func (a *API) RespondWithError(response http.ResponseWriter, code int, message string) {
	log.Errorf("An error occourred with a request - %d: %s", code, message)
	a.RespondWithJSON(response, code, []string{message})
}

func (a *API) RespondWithJSON(response http.ResponseWriter, code int, payload interface{}) {
	jsonData, _ := json.Marshal(payload)
	log.Printf("Success - %d", code)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write(jsonData)
}

func (a *API) RespondWithHTML(response http.ResponseWriter, code int, payload string) {
	log.Printf("Success - %d", code)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(code)
	response.Write([]byte(payload))
}

func (a *API) RespondWithImage(response http.ResponseWriter, code int, payload string, format string) {
	var contentType string
	if "png" == format {
		contentType = "image/png"
	} else {
		contentType = "image/jpeg"
	}
	image, _ := base64.StdEncoding.DecodeString(payload)

	log.Infof("%d, %s", code, http.StatusText(code))
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Type", contentType)
	response.WriteHeader(code)
	response.Write(image)
}
