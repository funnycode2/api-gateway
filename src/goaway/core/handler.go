package core

import "github.com/valyala/fasthttp"

//负责请求后端服务, Handler在filter中指定, 使用第一个非空Handler
//对于不同后端服务采用不同协议的情况, 也应当在Handler中处理
type Handler interface {
	Do(req *fasthttp.Request, res *fasthttp.Response)
}

type defaultHandler struct{}

func (h *defaultHandler) Do(req *fasthttp.Request, res *fasthttp.Response) {
	fasthttp.Do(req, res)
}
