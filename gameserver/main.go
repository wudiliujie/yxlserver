package main

import (
	"github.com/wudiliujie/common/module"
	"yxlserver/services/conf"
	"yxlserver/services/consts"
	"yxlserver/services/module/app"
	"yxlserver/services/module/gate"
)

func main() {
	conf.Server.ServerId = 1
	conf.Server.ServerType = consts.ServerType_Server
	conf.Server.ServerName = "网关"
	modules := []module.Module{app.CreateGameAppModule(), gate.Module}
	/*	if conf.Server.DebugAddr != "" {
		log.Release("启动debug端口:%v", conf.Server.DebugAddr)
		go http.ListenAndServe(conf.Server.DebugAddr, nil)
	}*/
	module.Run(modules...)
}
