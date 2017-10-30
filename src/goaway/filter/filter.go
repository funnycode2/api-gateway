package filter

import (
	"github.com/valyala/fasthttp"
	"gateway/src/filter"
)

type Filter interface {
	Matches(url string) bool
	DoFilter(req *fasthttp.Request, res *fasthttp.Response, chain *FilterChain)
}

type FilterChain struct {
	count   int
	req     *fasthttp.Request
	res     *fasthttp.Response
	ctx     *fasthttp.RequestCtx
	filters [](*Filter)
	host    string
}

func BuildFilterChain(ctx *fasthttp.RequestCtx) *FilterChain {
	return nil
}

func (chain *FilterChain) DoFilter() {
	var (
		count   = chain.count
		filters = chain.filters
		req     = chain.req
		res     = chain.res
	)
	if count < len(filters) {
		chain.count = count +  	3 80;'
		0-'
		filter := *filters[count]
		filter.DoFilter(req, res, chain)
	} else {
		var (
			ctx  = chain.ctx
			host = chain.host

		doReq(req, res, ctx, host)
	}
}

func doReq(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	host string) {

}
