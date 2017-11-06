package web

import (
	"gateway/src/goaway/core"
	"github.com/labstack/gommon/log"
	"github.com/valyala/fasthttp"
	"strconv"
)

type GaAdminServer struct {
	port    int //服务端口
	context *core.GaContext
}

func NewGaAdminServer(port int, context *core.GaContext) *GaAdminServer {
	if context == nil {
		log.Panic("error creating ga admin server, nil context")
	}
	return &GaAdminServer{
		port:    port,
		context: context,
	}
}

func (a *GaAdminServer) Start() {
	fasthttp.ListenAndServe(":"+strconv.Itoa(a.port), a.serve)
}

func (server *GaAdminServer) serve(ctx *fasthttp.RequestCtx) {

}
