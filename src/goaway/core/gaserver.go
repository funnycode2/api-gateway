package core

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"github.com/labstack/gommon/log"
	"sync"
)

type gaServer struct {
	port        int //监听端口号
	context     *context
	runOnce     *sync.Once  //启动方法只能调用一次
	contextLock *sync.Mutex //上下文加载控制锁
}

func NewGaServer(
	port int,
	context *context) *gaServer {
	return &gaServer{
		port:        port,
		context:     context,
		runOnce:     &sync.Once{},
		contextLock: &sync.Mutex{},
	}
}

//启动服务方法只能被调用一次
func (server *gaServer) Start() {
	server.runOnce.Do(server.start)
}

//提供上下文加载方法, 从而支持热配置
func (server *gaServer) LoadContext(c *context) {
	lock := server.contextLock
	if c == nil {
		log.Error("empty context, not loaded")
		return
	}
	lock.Lock()
	defer lock.Unlock()
	server.context = c
	log.Info("new context loaded")
}

func (server *gaServer) start() {
	log.Info("Ga-server listening on port: ", server.port)
	message := fasthttp.ListenAndServe(":"+strconv.Itoa(server.port), server.serve)
	log.Errorf("Ga-server (port: %d) exited due to error: %s", server.port, message)
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
