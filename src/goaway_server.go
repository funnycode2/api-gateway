package main

import (
	"gateway/src/goaway/core"
	"sync"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strings"
	"strconv"
	"gateway/src/goaway/constants"
	"gateway/src/goaway_example/server"
)

const (
	GOAWAY_PORT = 8888
	ADMIN_PORT  = 9999
)

var (
	appContext = server.NewMqlAppContext()
	reloadChan = make(chan int) //刷新配置管道
)

func main() {
	wg := &sync.WaitGroup{}
	gaContext := loadContext()
	gaServer := core.NewGaServer(GOAWAY_PORT, gaContext)
	//LoadContext用于热配置
	wg.Add(1)
	//网关服务
	go func() {
		gaServer.Start()
		wg.Done()
	}()
	//web管理界面
	go func() {
		startWebServer()
	}()
	//重新载入配置
	go func() {
		for {
			signal := <-reloadChan
			if signal == 1 {
				gaServer.LoadContext(loadContext())
			}
		}
	}()
	wg.Wait()
}

func startWebServer() {
	context := core.NewContext()
	context.LoadFilter(&server.CORSFilter{})
	context.LoadFilter(&server.StaticFileFilter{})
	context.LoadFilter(&ServiceFilter{})
	webServer := core.NewGaServer(ADMIN_PORT, context)
	webServer.Start()
}

func loadContext() *core.GaContext {
	gaContext := core.NewContext()
	appContext.VisitUriHosts(gaContext)
	appContext.VisitUriFilters(gaContext)
	return gaContext
}

//核心服务代码实现
type ServiceFilter struct {
	core.BaseFilter
}

func (b *ServiceFilter) Matches(uri string) bool {
	return true
}

func (b *ServiceFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	uri := string(req.Header.RequestURI())
	if strings.HasPrefix(uri, "/admin/server/list") {
		prefix := string(ctx.QueryArgs().Peek("prefix"))
		desc := string(ctx.QueryArgs().Peek("desc"))
		currentpage := string(ctx.QueryArgs().Peek("currentpage"))
		currentpageInt := 0
		if len(currentpage) > 0 {
			currentpageInt, _ = strconv.Atoi(currentpage)
		}
		result := appContext.QueryService(prefix, desc, currentpageInt)
		res.Header.SetBytesKV(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
		jsonStr, _ := json.Marshal(*result)
		res.SetBody(jsonStr)
	}
	if strings.HasPrefix(uri, "/admin/server/modify") {
		var service server.Mservice
		json.Unmarshal(req.Body(), &service)
		err := appContext.UpdateService(&service)
		if err != nil {
			res.AppendBodyString("0")
		} else {
			res.AppendBodyString("1")
		}
	}
	if strings.HasPrefix(uri, "/admin/server/reload") {
		reloadChan <- 1
		res.AppendBodyString("1")
	}
}
