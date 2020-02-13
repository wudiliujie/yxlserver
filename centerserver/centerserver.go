package main

import (
	"github.com/wudiliujie/common/module"
	"yxlserver/services/conf"
	"yxlserver/services/consts"
	"yxlserver/services/module/app"
	"yxlserver/services/module/gate"
)

func main() {
	conf.Server.WSAddr = "0.0.0.0:8888"
	conf.Server.ServerType = consts.ServerType_Center
	modules := []module.Module{gate.Module, app.CreateCenterAppModule()}
	/*	if conf.Server.DebugAddr != "" {
		log.Release("启动debug端口:%v", conf.Server.DebugAddr)
		go http.ListenAndServe(conf.Server.DebugAddr, nil)
	}*/
	module.Run(modules...)
}
