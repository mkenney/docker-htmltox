/*
Package main contains the HTML Conversion Service that interfaces with Chrome
*/
package main

import (
	"fmt"
	"log"

	chrome "github.com/mkenney/go-chrome"
)

func main() {
	browser, err := chrome.GetBrowser()
	if nil != err {
		log.Fatalf("Could not load browser instance: %v", err)
		return
	}
	log.Printf("Browser instance initialized")
	tabs, _ := browser.GetTabs()
	for _, tab := range tabs {
		fmt.Printf("Tab: %v\n", tab)
		fmt.Printf("\tDescription: %v\n", tab.Description)
		fmt.Printf("\tDevtoolsFrontendURL: %v\n", tab.DevtoolsFrontendURL)
		fmt.Printf("\tID: %v\n", tab.ID)
		fmt.Printf("\tTitle: %v\n", tab.Title)
		fmt.Printf("\tType: %v\n", tab.Type)
		fmt.Printf("\tURL: %v\n", tab.URL)
		fmt.Printf("\tWebSocketDebuggerURL: %v\n", tab.WebSocketDebuggerURL)
		fmt.Printf("\tSocket: %v\n", tab.Socket)
	}

	browser.Close()
}

/*
New returns a pointer to an HTMLToX struct
*/
func New() *HTMLToX {
	htmltox := new(HTMLToX)
	return htmltox
}

/*
Render takes an HTML source, either a string or a URL, and returns
a byte array of the resulting image

@param source An HTML string or URL
@param format An output format, one of 'jpg', 'png', 'pdf'
@param width The viewport width
@param height The viewport height
*/
func (htmltox *HTMLToX) Render(source, format string, width, height int) (result []byte) {

	return result
}

/*
HTMLToX defines the struct for the HTML conversion API service
*/
type HTMLToX struct{}
