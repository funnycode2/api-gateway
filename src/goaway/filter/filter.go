package filter

import (
	"github.com/valyala/fasthttp"
	"gateway/src/goaway/mapping"
	"gateway/src/goaway/handler"
)

type Filter interface {
	Matches(url string) bool
	DoFilter(req *fasthttp.Request, res *fasthttp.Response, chain *filterChain)
}

type filterChain struct {
	count   int
	req     *fasthttp.Request
	res     *fasthttp.Response
	ctx     *fasthttp.RequestCtx
	filters []Filter
	mapping mapping.Mapping
	handler handler.Handler
}

func BuildFilterChain(
	ctx *fasthttp.RequestCtx,
	filters []Filter,
	mapping mapping.Mapping,
	handler handler.Handler) *filterChain {
	return &filterChain{
		count:   0,
		req:     &ctx.Request,
		res:     &ctx.Response,
		ctx:     ctx,
		filters: filters,
		mapping: mapping,
		handler: handler,
	}
}

func (chain *filterChain) DoFilter() {
	var (
		count   = chain.count
		filters = chain.filters
		req     = chain.req
		res     = chain.res
		mapping = chain.mapping
		ctx     = chain.ctx
		handler = chain.handler
	)
	if count < len(filters) {
		chain.count = count + 1
		filters[count].DoFilter(req, res, chain)
	} else {
		if mapping != nil {
			target := mapping.TargetHost()
			handler.Handle(req, res, ctx, target)
		}
	}
}
