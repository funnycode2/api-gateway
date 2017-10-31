package core

import "github.com/valyala/fasthttp"

//负责请求后端服务
type Handler interface {
	Do(req *fasthttp.Request, res *fasthttp.Response)
}

type defaultHandler struct{}

func (h *defaultHandler) Do(req *fasthttp.Request, res *fasthttp.Response) {
	fasthttp.Do(req, res)
}
