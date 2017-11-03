package core

import (
	"github.com/valyala/fasthttp"
	"github.com/labstack/gommon/log"
)

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

	//使用原始请求作为后端服务请求, 原因是在文件上传的时候
	//通过CopyTo方法得到的请求会丢失请求体, 其具体原因还有待确认
	upRes := fasthttp.AcquireResponse()
	upRes.Reset()
	defer fasthttp.ReleaseResponse(upRes)

	chain.DoFilter(req, upRes, ctx)

	//将后端服务的响应最终写回
	upRes.Header.CopyTo(&res.Header)
	res.AppendBody(upRes.Body())
}

var gaDefaultHandler = &defaultHandler{}

func (f *coreFilter) Handler() Handler {
	return gaDefaultHandler
}

func logError() {
	if err := recover(); err != nil {
		log.Error("error occurred: ", err)
	}
}
