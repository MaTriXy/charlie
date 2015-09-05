package main

import (
	"log"
	"net/http"

	"github.com/f2prateek/charlie/ci"
	"github.com/tj/docopt"
)

const (
	Version = "1.0.0"
	Usage   = `Charlie.

Simple CI for projects built with Gradle.

Usage:
  charlie [--port=<port>]
  charlie -h | --help
  charlie --version

Options:
  --port=<port>    port [default: 3001].
  -h --help        Show this screen.
  --version        Show version.`
)

func main() {
	arguments, _ := docopt.Parse(Usage, nil, true, Version, false)

	port := arguments["--port"].(string)

	ci := &ci.Server{}
	check(http.ListenAndServe(":"+port, ci))
}

func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
