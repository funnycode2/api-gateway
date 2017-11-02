package main

import (
	"gateway/src/goaway/core"
	"gateway/src/goaway/ext"
)

const (
	port = 8888
)

func main() {
	context := core.NewContext()
	//context.AddFilter(&ext.OauthFilter{})
	filter := &ext.GapFilter{}
	context.AddFilter(filter)
	gaServer := core.NewGaServer(port, context)
	gaServer.Start()
}