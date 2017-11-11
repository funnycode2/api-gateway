package main

import (
	ex "gateway/src/goaway_example"
	"gateway/src/goaway/core"
	"sync"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strings"
	"strconv"
	"gateway/src/goaway/constants"
	"gateway/src/goaway_example/web"
)

const (
	port = 8888
)

var (
	appContext = ex.NewMqlAppContext()
)

func main() {
	wg := &sync.WaitGroup{}
	gaContext := core.NewContext()
	appContext := ex.NewMqlAppContext()
	appContext.VisitUriHosts(gaContext)
	appContext.VisitUriFilters(gaContext)
	gaServer := core.NewGaServer(port, gaContext)
	//LoadContext用于热配置
	wg.Add(1)
	go func() {
		gaServer.Start()
		wg.Done()
	}()
	go func() {
		startWebServer()
	}()
	wg.Wait()
}

func startWebServer() {
	context := core.NewContext()
	context.LoadFilter(&ex.CORSFilter{})
	context.LoadFilter(&StaticFileFilter{})
	context.LoadFilter(&JsonFilter{})
	server := core.NewGaServer(9999, context)
	server.Start()
}

type JsonFilter struct {
	core.BaseFilter
}

func (b *JsonFilter) Matches(uri string) bool {
	return true
}

func (b *JsonFilter) DoFilter(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	chain *core.FilterChain) {
	uri := string(req.Header.RequestURI())
	if strings.HasPrefix(uri, "/admin/service/list") {
		prefix := string(ctx.QueryArgs().Peek("prefix"))
		desc := string(ctx.QueryArgs().Peek("desc"))
		currentpage := string(ctx.QueryArgs().Peek("currentpage"))
		currentpageInt := 0
		if len(currentpage) > 0 {
			currentpageInt, _ = strconv.Atoi(currentpage)
		}
		result := appContext.QueryService(prefix, desc, currentpageInt)
		res.Header.SetBytesKV(constants.CONTENT_TYPE, []byte("application/json"))
		json, _ := json.Marshal(*result)
		res.SetBody(json)
	}
	if strings.HasPrefix(uri, "/admin/service/modify") {
		var service web.Mservice
		json.Unmarshal(req.Body(), &service)
		err := appContext.UpdateService(&service)
		if err != nil {
			res.AppendBodyString("0")
		} else {
			res.AppendBodyString("1")
		}
	}
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
