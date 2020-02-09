package main

import (
	"github.com/wudiliujie/common/module"

	"yxlserver/services/gateservice/module/gate"
)

func main() {
	modules := []module.Module{gate.Module}
	/*	if conf.Server.DebugAddr != "" {
		log.Release("启动debug端口:%v", conf.Server.DebugAddr)
		go http.ListenAndServe(conf.Server.DebugAddr, nil)
	}*/
	module.Run(modules...)
}
