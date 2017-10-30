package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"gateway/src/goaway/core"
)

const (
	port = ":8888"
)

func main() {
	log.Printf("listening on localhost%s\n", port)
	log.Printf("Proxy exit at %s", fasthttp.ListenAndServe(port, core.HttpHandler))
}