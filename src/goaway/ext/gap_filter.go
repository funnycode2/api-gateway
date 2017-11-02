package ext

import (
	"github.com/valyala/fasthttp"
	"strings"
	"gateway/src/goaway/core"
)

type GapFilter struct{}

func (f *GapFilter) Matches(url string) bool {
	return strings.HasPrefix(url, "/gap")
}

func (f *GapFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	req.SetHost("localhost:8080")
	println("up request:")
	println(req.String())
	chain.DoFilter(req, res, ctx)
	println("up response:")
	println(res.String())
}
