package core

import (
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/context"
	"gateway/src/goaway/filter"
)

func HttpHandler(ctx *fasthttp.RequestCtx) {
	var (
		appctx  = context.Context
		uri     = (string)(ctx.Request.Header.RequestURI())
		filters = appctx.Filters
	)

	//将匹配的过滤器找出来, 按顺序组成数组 (核心过滤器(CoreFilter)总是在第一个)
	var matchFilters []filter.Filter
	for _, f := range filters {
		match := f.Matches(uri)
		if match {
			matchFilters = append(matchFilters, f)
		}
	}

	filterChain := filter.NewFilterChain(matchFilters)
	filterChain.DoFilter(&ctx.Request, &ctx.Response, ctx, filterChain)
}
