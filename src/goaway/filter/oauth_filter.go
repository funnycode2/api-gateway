package filter

import (
	"github.com/valyala/fasthttp"
	"github.com/labstack/gommon/log"
)

type OauthFilter struct{}

func (f *OauthFilter) Matches(url string) bool {
	return true
}

func (f *OauthFilter) DoFilter(req *fasthttp.Request, res *fasthttp.Response, chain *filterChain) {
	log.Info("entering oauth filter...")
	chain.DoFilter()
	log.Info("exiting oauth filter...")
}
