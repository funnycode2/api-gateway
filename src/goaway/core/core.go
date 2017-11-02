package core

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"github.com/labstack/gommon/log"
)

type gaServer struct {
	port    int //监听端口号
	context *context
}

func NewGaServer(
	port int,
	context *context) *gaServer {
	return &gaServer{
		port:    port,
		context: context,
	}
}

func (server *gaServer) Start() {
	log.Info("Ga-server listening on port: ", server.port)
	message := fasthttp.ListenAndServe(":"+strconv.Itoa(server.port), server.serve)
	log.Errorf("Ga-server (port: %d) exited due to error: 	%s", server.port, message)
}

func (server *gaServer) serve(ctx *fasthttp.RequestCtx) {
	var (
		gaCtx   = server.context
		uri     = string(ctx.Request.Header.RequestURI())
		filters = gaCtx.Filters()
	)

	//将匹配的过滤器找出来, 按顺序组成数组 (核心过滤器(gaCoreFilter)总是在第一个)
	var matchFilters []Filter
	for _, f := range *filters {
		match := f.Matches(uri)
		if match {
			matchFilters = append(matchFilters, f)
		}
	}

	filterChain := NewFilterChain(matchFilters)
	filterChain.DoFilter(&ctx.Request, &ctx.Response, ctx)
}
