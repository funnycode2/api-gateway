package core

import (
	"github.com/valyala/fasthttp"
	"github.com/labstack/gommon/log"
)

type CoreFilter struct{}

//系统核心过滤器, 该过滤器必须作为第一个过滤器来执行
func (f *CoreFilter) Matches(url string) bool {
	return true
}

func (f *CoreFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *FilterChain) {

	//记录错误信息
	defer logError()

	upReq := fasthttp.AcquireRequest()
	upReq.Reset()
	req.CopyTo(upReq)
	defer fasthttp.ReleaseRequest(upReq)

	upRes := fasthttp.AcquireResponse()
	upRes.Reset()
	defer fasthttp.ReleaseResponse(upRes)

	chain.DoFilter(upReq, upRes, ctx, chain)

	//将后端服务的响应最终写回
	upRes.Header.CopyTo(&res.Header)
	res.AppendBody(upRes.Body())
}

func logError() {
	if err := recover(); err != nil {
		log.Error("error occured: ", err)
	}
}