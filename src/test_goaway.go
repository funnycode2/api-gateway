package main

import (
	ex "gateway/src/goaway_example"
	"gateway/src/goaway/core"
	"sync"
)

const (
	port = 8888
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
	wg.Wait()
}
