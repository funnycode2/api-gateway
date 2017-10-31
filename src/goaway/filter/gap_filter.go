package filter

import (
	"github.com/valyala/fasthttp"
	"strings"
)

type GapFilter struct{}

func (GapFilter) Matches(url string) bool {
	return strings.HasPrefix(url, "/gap")
}

func (GapFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *filterChain) {
	req.SetHost("localhost:8080")
}


