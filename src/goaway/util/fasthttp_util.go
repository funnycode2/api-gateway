package util

import "github.com/valyala/fasthttp"

func HttpCopy(req *fasthttp.Request) *fasthttp.Request {
	newreq := fasthttp.AcquireRequest()
	newreq.Reset()
	req.CopyTo(newreq)
	return newreq
}