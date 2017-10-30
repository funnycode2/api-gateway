package filter

import (
	"github.com/valyala/fasthttp"
	"github.com/labstack/gommon/log"
)

type DefaultFilter struct{}

func (f *DefaultFilter) Matches(url string) bool {
	return true
}

func (f *DefaultFilter) DoFilter(req *fasthttp.Request, res *fasthttp.Response, chain *filterChain) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic in processing! ", err)
		}
		//req.Reset()
		//res.Reset()
	}()
	chain.DoFilter()
}
