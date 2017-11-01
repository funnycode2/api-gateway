package core

import (
	"github.com/valyala/fasthttp"
	"github.com/labstack/gommon/log"
)

type Filter interface {
	Matches(url string) bool

	DoFilter(
		req *fasthttp.Request,
		res *fasthttp.Response,
		ctx *fasthttp.RequestCtx,
		chain *FilterChain)
}

//系统核心过滤器, 该过滤器必须作为第一个过滤器来执行
type coreFilter struct{}

func (f *coreFilter) Matches(url string) bool {
	return true
}

func (f *coreFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *FilterChain) {
	//捕获并记录错误信息
	defer logError()

	//使用原始请求作为后端服务请求, 原因是在文件上传的时候通过 CopyTo方法得到的请求会丢失请求体, 其具体原因还有待确认
	/*upReq := fasthttp.AcquireRequest()
	upReq.Reset()
	req.CopyTo(upReq)
	defer fasthttp.ReleaseRequest(upReq)*/

	upRes := fasthttp.AcquireResponse()
	upRes.Reset()
	defer fasthttp.ReleaseResponse(upRes)

	chain.DoFilter(req, upRes, ctx, chain)
	//chain.DoFilter(upReq, upRes, ctx, chain)

	//将后端服务的响应最终写回
	upRes.Header.CopyTo(&res.Header)
	res.AppendBody(upRes.Body())
}

func logError() {
	if err := recover(); err != nil {
		log.Error("error occurred: ", err)
	}
}
