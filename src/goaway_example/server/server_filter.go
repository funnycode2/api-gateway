package server

import (
	"gateway/src/goaway/core"
	"github.com/valyala/fasthttp"
	"strings"
	"gateway/src/goaway/util"
	"github.com/labstack/gommon/log"
	"gateway/src/goaway/constants"
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

type StaticFileFilter struct {
	core.BaseFilter
}

func (b *StaticFileFilter) Matches(uri string) bool {
	return strings.HasSuffix(uri, ".html") ||
		strings.HasSuffix(uri, ".js") ||
		strings.HasSuffix(uri, ".css") ||
		strings.HasSuffix(uri, ".eot") ||
		strings.HasSuffix(uri, ".ttf") ||
		strings.HasSuffix(uri, ".woff2") ||
		strings.HasSuffix(uri, ".woff")
}

func (b *StaticFileFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	uri := string(req.Header.RequestURI())
	if strings.HasSuffix(uri, ".ttf") {
		res.Header.Set("Accept-Ranges", "bytes")
		res.Header.SetBytesKV(constants.CONTENT_TYPE, []byte("font/ttf"))
	} else if strings.HasSuffix(uri, ".woff2") {
		res.Header.Set("Accept-Ranges", "bytes")
		res.Header.SetBytesKV(constants.CONTENT_TYPE, []byte("font/woff2"))
	} else {
		accept := string(req.Header.Peek("Accept"))
		contentType := strings.Split(accept, ",")[0]
		res.Header.SetBytesKV(constants.CONTENT_TYPE, []byte(contentType+";UTF-8"))
	}
	res.SendFile("src/goaway_example/web/" + uri)
}