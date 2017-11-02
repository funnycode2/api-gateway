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
	context.AddFilter(&ext.GapFilter{})
	gaServer := core.NewGaServer(port, context)
	gaServer.Start()
}