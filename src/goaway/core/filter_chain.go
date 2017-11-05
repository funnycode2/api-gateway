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

func newFilterChain(
	filters []Filter) *FilterChain {
	//默认请求处理器为核心过滤器的
	handler := gaCoreFilter.Handler()
	if len(filters) > 0 {
		//如果过滤器中有定义请求处理器, 则使用最先一个过滤器中定义的请求处理器
		for _, filter := range filters {
			if filter.Handler() != nil {
				handler = filter.Handler()
				break
			}
		}
	}
	//核心过滤器总是在第一个
	filters = append([]Filter{gaCoreFilter}, filters...)
	return &FilterChain{
		count:   0,
		filters: filters,
		handler: handler,
	}
}

func (chain *FilterChain) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx) {

	var (
		count   = chain.count
		filters = chain.filters
	)

	if count < len(filters) {
		chain.count = count + 1
		filters[count].DoFilter(req, res, ctx, chain)
	} else {
		//最后一个过滤器执行完后会向后端请求服务
		chain.handler.Do(req, res)
	}
}
