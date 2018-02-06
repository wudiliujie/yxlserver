package main

import (
	"common"
	"common/msg"
	"leaf"
	lconf "leaf/conf"
	"gameclient/client"
)

func main()  {
	common.Init()
	client.Init(msg.Processor)
	lconf.ConsolePort=9999
	leaf.Run(
		client.Module,
	)
}
