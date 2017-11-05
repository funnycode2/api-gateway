package goaway_example

import (
	"gateway/src/goaway/core"
	"github.com/valyala/fasthttp"
	"strings"
	"gateway/src/goaway/util"
	"github.com/labstack/gommon/log"
)

//适配当前的mysql数据库配置, 功能是重写目标主机(含端口)
type forwardFilter struct {
	core.BaseFilter
	uri  string //匹配的URI
	host string //目标主机含端口
}

func NewForwardFilter(uri string, host string) *forwardFilter {
	normalizedUri, err := util.NormalizeUri(uri)
	if err != nil {
		log.Panicf("Invalid Uri: %s", uri)
	}
	if util.MatchHost(host) {
		log.Printf("Invalid host: %s", host)
	}
	return &forwardFilter{
		uri:  normalizedUri,
		host: host,
	}
}

func (a *forwardFilter) Matches(uri string) bool {
	return strings.HasPrefix(uri, a.uri)
}

func (a *forwardFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	req.SetHost(a.host)
	log.Print("forward request:\n", req)
	chain.DoFilter(req, res, ctx)
}

func (a *forwardFilter) String() string {
	return "Uri: " + a.uri + ", host: " + a.host
}
