/*
Package main executes the HTML Conversion Service
*/
package main

import (
	htmltox "github.com/mkenney/docker-htmltox/app/htmltox"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&textFormat{})
}

func main() {
	htmltox, err := htmltox.New()
	if nil != err {
		log.Fatalf("Could not initialize conversion service: %s", err.Error())
	}
	log.Info("Starting API server")
	htmltox.API.Run(80)
}
