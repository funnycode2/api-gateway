package core

import (
	"github.com/valyala/fasthttp"
)

//简单的责任链模式
type FilterChain struct {
	count   int
	filters []Filter
	handler Handler
}

var gaDefaultHandler = &defaultHandler{}

func NewFilterChain(
	filters []Filter) *FilterChain {
	return &FilterChain{
		count:   0,
		filters: filters,
	}
}

func (chain *FilterChain) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain2 *FilterChain) {

	var (
		count   = chain.count
		filters = chain.filters
	)

	if count < len(filters) {
		chain.count = count + 1
		filters[count].DoFilter(req, res, ctx, chain)
	} else {
		//最后一个过滤器执行完后会向后端请求服务
		handler := chain.handler
		if handler == nil {
			gaDefaultHandler.Do(req, res)
		} else {
			handler.Do(req, res)
		}
	}

}

func (chain *FilterChain) SetHandler(handler *Handler) {
	if handler != nil {
		chain.handler = handler
	}
}
