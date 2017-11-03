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
	//context.AddFilter(&ext.OauthFilter{})
	context.AddFilter(ext.NewBasicServiceFilter(
		"/gap", "/gap", "newyuncaijia.igap.cc"))
	gaServer := core.NewGaServer(port, context)
	//LoadContext用于热配置
	gaServer.LoadContext(nil)
	go func() {
		gaServer.Start()
		wg.Done()
	}()
	wg.Wait()
}
