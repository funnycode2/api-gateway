package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"gateway/src/goaway/core"
	"gateway/src/goaway/ext"
)

const (
	port = ":8888"
)

func main() {

	context := core.Context{}
	//context.AddFilter(&ext.OauthFilter{})
	context.AddFilter(&ext.GapFilter{})
	gaServer := core.NewGaServer(&context)

	log.Printf("listening on localhost%s\n", port)
	log.Printf("Proxy exit at %s", fasthttp.ListenAndServe(port, gaServer.Serve))
}