package main

import (
	"gateway/src/goaway/core"
	"gateway/src/goaway/ext"
	"sync"
)

const (
	port = 8888
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	context := core.NewContext()
	//context.LoadFilter(&ext.OauthFilter{})
	context.LoadFilter(ext.NewBasicServiceFilter(
		"/1", "/1", "news.mydrivers.com"))
	gaServer := core.NewGaServer(port, context)
	//LoadContext用于热配置
	gaServer.LoadContext(context)
	go func() {
		gaServer.Start()
		wg.Done()
	}()
	wg.Wait()
}
