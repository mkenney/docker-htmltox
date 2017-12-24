/*
Package main executes the HTML Conversion Service
*/
package main

import (
	htmltox "github.com/mkenney/docker-htmltox/app/htmltox"
	log "github.com/sirupsen/logrus"
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
