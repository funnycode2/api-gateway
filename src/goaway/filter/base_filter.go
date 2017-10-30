package filter

import "github.com/valyala/fasthttp"

var (
	BASE_FILTER = &baseFilter{}
)

type baseFilter struct {
	Filter
}

func (f *baseFilter) Matches(url string) bool {
	return true
}

func (f *baseFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	chain *FilterChain) {
	defer func() {
		req.Reset()
		res.Reset()
	}()
	chain.DoFilter()
}
