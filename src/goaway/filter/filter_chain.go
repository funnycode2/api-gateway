package filter

import (
	"github.com/valyala/fasthttp"
)

//简单的责任链模式
type filterChain struct {
	count   int
	filters []Filter
}

func NewFilterChain(
	filters []Filter) *filterChain {
	return &filterChain{
		count:   0,
		filters: filters,
	}
}

func (chain *filterChain) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain2 *filterChain) {

	var (
		count   = chain.count
		filters = chain.filters
	)

	total := len(filters) - 1
	if count <= total {
		chain.count = count + 1
		filters[count].DoFilter(req, res, ctx, chain)
	}

	//最后一个过滤器执行完后会向后端请求服务
	if count == total {
		fasthttp.Do(req, res)
	}
}