package core

import (
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/context"
	ghandler "gateway/src/goaway/mapping"
	"gateway/src/goaway/filter"
	"gateway/src/goaway/handler"
)

func HttpHandler(ctx *fasthttp.RequestCtx) {
	var (
		appctx   = context.Context
		req      = ctx.Request
		uri      = (string)(req.Header.RequestURI())
		filters  = appctx.Filters
		mappings = appctx.Mappings
		handlers = appctx.Handlers
	)

	var matchMapping ghandler.Mapping
	for _, h := range mappings {
		match := h.Matches(uri)
		if match {
			matchMapping = h
			break
		}
	}

	if mappings == nil {
		return
	}

	matchFilters := make([]filter.Filter, 0)
	for _, f := range filters {
		match := f.Matches(uri)
		if match {
			matchFilters = append(matchFilters, f)
		}
	}

	var matchHandler handler.Handler
	for _, h := range handlers {
		match := h.Matches(uri)
		if match {
			matchHandler = h
			break
		}
	}

	if matchHandler == nil {
		return
	}

	filterChain := filter.BuildFilterChain(ctx, matchFilters, matchMapping, matchHandler)
	filterChain.DoFilter()
}
