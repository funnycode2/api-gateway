package main

import (
	"github.com/valyala/fasthttp"
	"log"
)

const (
	test_port = ":9999"
)

func main() {
	log.Printf("listening on localhost%s\n", test_port)
	log.Printf("Proxy exit at %s", fasthttp.ListenAndServe(test_port, func (ctx *fasthttp.RequestCtx) {
		ctx.Response.AppendBodyString("hello goaway")
	}))
}