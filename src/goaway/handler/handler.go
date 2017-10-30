package handler

import "github.com/valyala/fasthttp"

type Handler interface {
	Matches(url string) bool

	Handle(
		req *fasthttp.Request,
		res *fasthttp.Response,
		ctx *fasthttp.RequestCtx,
		mapping string)
}