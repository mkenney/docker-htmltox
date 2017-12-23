/*
Package main executes the HTML Conversion Service
*/
package main

import (
	log "github.com/Sirupsen/logrus"
	htmltox "github.com/mkenney/docker-htmltox/app/htmltox"
)

func main() {
	log.Info("Starting Chrome")
	htmltox, err := htmltox.New()
	if nil != err {
		log.Fatalf("Could not initialize conversion service: %v", err)
	}
	log.Info("Starting API server")
	htmltox.API.Run(80)
}

//	for {
//		break
//	}
//
//	// browser testing
//	var browser *chrome.Browser
//	var conn *chrome.Socket
//	browser, err = chrome.GetBrowser()
//	log.Infof("Browser instance initialized")
//	tabs, _ := browser.GetTabs()
//	for _, tab := range tabs {
//		fmt.Printf("Tab: %v\n", tab)
//		fmt.Printf("\tDescription: %v\n", tab.Description)
//		fmt.Printf("\tDevtoolsFrontendURL: %v\n", tab.DevtoolsFrontendURL)
//		fmt.Printf("\tID: %v\n", tab.ID)
//		fmt.Printf("\tTitle: %v\n", tab.Title)
//		fmt.Printf("\tType: %v\n", tab.Type)
//		fmt.Printf("\tURL: %v\n", tab.URL)
//		fmt.Printf("\tWebSocketDebuggerURL: %v\n", tab.WebSocketDebuggerURL)
//		fmt.Printf("\tSocket: %v\n", tab.Socket)
//	}
//
//	conn, err = browser.NewBrowserSocket()
//	log.Infof("Websocket: %v\n", conn)
//
//	browser.Close()
//}
//
