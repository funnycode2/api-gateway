package core

import (
	"github.com/valyala/fasthttp"
)

type Filter interface {
	Matches(url string) bool

	DoFilter(
		req *fasthttp.Request,
		res *fasthttp.Response,
		ctx *fasthttp.RequestCtx,
		chain *FilterChain)
}