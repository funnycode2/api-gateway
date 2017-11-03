package core

import (
	"github.com/valyala/fasthttp"
)

type Filter interface {
	Matches(url string) bool

	DoFilter(
		req *fasthttp.Request,
		res *fasthttp.Response,
		ctx *fasthttp.RequestCtx,
		chain *FilterChain)

	Handler() Handler

	OnDestroy()
}

type BaseFilter struct {
}

func (b *BaseFilter) Matches(url string) bool {
	return true
}

func (b *BaseFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *FilterChain) {
	chain.DoFilter(req, res, ctx)
}

func (b *BaseFilter) Handler() Handler {
	return nil
}

func (b *BaseFilter) OnDestroy() {}
