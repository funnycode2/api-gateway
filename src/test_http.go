package main

import (
	"github.com/valyala/fasthttp"
)

func main() {
	statusCode, body, err := fasthttp.Get(nil, "http://localhost:9999/")
	println(statusCode)
	println(string(body))
	if err != nil {
		println(err)
	}
}
