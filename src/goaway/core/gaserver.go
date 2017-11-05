package core

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"github.com/labstack/gommon/log"
	"sync"
)

type gaServer struct {
	port        int //监听端口号
	context     *GaContext
	runOnce     *sync.Once  //启动方法只能调用一次
	contextLock *sync.Mutex //上下文加载控制锁
}

func NewGaServer(
	port int,
	context *GaContext) *gaServer {
	if port < 0 {
		log.Panic("Invalid port number, must be positive number")
	}
	if context == nil {
		log.Panic("Nil GaContext not allowed")
	}
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
func (server *gaServer) LoadContext(c *GaContext) {
	if c == nil {
		log.Error("empty GaContext, not loaded")
		return
	}
	lock := server.contextLock
	lock.Lock()
	defer lock.Unlock()
	go server.context.onDestroy() //为了快速切换上下文
	server.context = c
	log.Info("new GaContext loaded")
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
		filters = &gaCtx.filters
	)

	//将匹配的过滤器找出来, 按顺序组成数组
	var matchFilters []Filter
	for _, f := range *filters {
		match := f.Matches(uri)
		if match {
			matchFilters = append(matchFilters, f)
		}
	}

	filterChain := newFilterChain(matchFilters)
	filterChain.DoFilter(&ctx.Request, &ctx.Response, ctx)
}
