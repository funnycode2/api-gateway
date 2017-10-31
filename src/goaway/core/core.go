package core

import (
	"github.com/valyala/fasthttp"
)

type GaServer struct {
	context *Context
}
var (
	coreFilter = &CoreFilter{}
)

func NewGaServer(context *Context) *GaServer {
	var gaFilter []Filter
	filters := context.Filters
	if filters == nil || len(filters) == 0 {
		gaFilter = []Filter{coreFilter}
	} else {
		if filters[0] != coreFilter {
			gaFilter = append([]Filter{coreFilter}, filters...)
		}
	}
	context.Filters = gaFilter
	return &GaServer{
		context: context,
	}
}

func (server*GaServer) Serve(ctx *fasthttp.RequestCtx) {
	var (
		gaCtx   = server.context
		uri     = string(ctx.Request.Header.RequestURI())
		filters = gaCtx.Filters
	)

	//将匹配的过滤器找出来, 按顺序组成数组 (核心过滤器(CoreFilter)总是在第一个)
	var matchFilters []Filter
	for _, f := range filters {
		match := f.Matches(uri)
		if match {
			matchFilters = append(matchFilters, f)
		}
	}

	filterChain := NewFilterChain(matchFilters)
	filterChain.DoFilter(&ctx.Request, &ctx.Response, ctx, filterChain)
}
