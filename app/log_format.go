/*
Package main executes the Secure Proxy Service

This file defines a custom log format in the form of:
[ISO-8601-date] [hostname] [level] [caller] - [message]
*/
package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

/*
Set the formatter and the default level
The level is defined by the LOG_LEVEL environment variable. Default is 'info'
*/
func init() {
	levelFlag := os.Getenv("LOG_LEVEL")
	if "" == levelFlag {
		levelFlag = "info"
	}

	level, err := log.ParseLevel(levelFlag)
	if nil != err {
		log.Fatalf("Could not parse log level flag: %s", err)
	}

	formatter := new(logFormat)
	log.SetFormatter(formatter)
	log.SetLevel(level)
}

var logTemplate = template.Must(template.New("log").Funcs(funcMap).Parse("{{.Timestamp}} {{.Hostname}} {{Pad .Level 4 | Upper}} {{.Caller}} {{.Message}}"))

var funcMap = template.FuncMap{
	"Upper": strings.ToUpper,
	"Pad":   pad,
}

func pad(str string, length int) string {
	for {
		str += " "
		if len(str) > length {
			return str[0:length]
		}
	}
}

type logFormat struct {
	*log.TextFormatter
}

func (l *logFormat) Format(entry *log.Entry) ([]byte, error) {
	var logLine *bytes.Buffer
	RFC3339Milli := "2006-01-02T15:04:05.000Z07:00"

	if entry.Buffer != nil {
		logLine = entry.Buffer
	} else {
		logLine = &bytes.Buffer{}
	}

	caller := ""
	a := 0
	for {
		if pc, file, line, ok := runtime.Caller(a + 1); ok {
			if !strings.Contains(file, "github.com/Sirupsen/logrus") {
				caller = fmt.Sprintf("%s:%d %s -", path.Base(file), line, runtime.FuncForPC(pc).Name())
				break
			}
		} else {
			break
		}
		a++
	}

	data := struct {
		Timestamp string
		Hostname  string
		Level     string
		Caller    string
		Message   string
	}{
		entry.Time.Format(RFC3339Milli),
		os.Getenv("HOSTNAME"),
		entry.Level.String(),
		caller,
		entry.Message,
	}
	logTemplate.Execute(logLine, data)
	logLine.WriteByte('\n')
	return logLine.Bytes(), nil
}
