package ext

import (
	"github.com/valyala/fasthttp"
	"strings"
	"gateway/src/goaway/core"
)

type GapFilter struct{}

func (GapFilter) Matches(url string) bool {
	return strings.HasPrefix(url, "/gap")
}

func (GapFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	req.SetHost("localhost:8080")
	uri := string(req.RequestURI())
	println(uri)
	if strings.HasPrefix(uri, "/gap/login") {
		req.SetRequestURI("/gap/login")
	}
	println(req.String())
	chain.DoFilter(req, res, ctx, chain)
	println(res.String())
}
