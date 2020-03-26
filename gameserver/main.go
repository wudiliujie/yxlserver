package main

import (
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"github.com/wudiliujie/common/mysql"
	"time"
	"yxlserver/services/conf"
	"yxlserver/services/consts"
	"yxlserver/services/logic/reload"
	"yxlserver/services/module/app"
	"yxlserver/services/module/db"
	"yxlserver/services/module/gate"
)

func main() {
	//Test()
	conf.Server.ServerId = 1
	conf.Server.ServerType = consts.ServerType_Server
	conf.Server.ServerName = "网关"
	mysql.Init("root:root@tcp(192.168.22.203:3306)/gamedb?charset=utf8&parseTime=true&loc=Local", 100)
	reload.Reload()
	modules := []module.Module{app.CreateGameAppModule(), gate.Module, db.Module}
	/*	if conf.Server.DebugAddr != "" {
		log.Release("启动debug端口:%v", conf.Server.DebugAddr)
		go http.ListenAndServe(conf.Server.DebugAddr, nil)
	}*/

	module.Run(modules...)
}
func Test() {
	intAttr := make(map[int]interface{})
	for i := 0; i < 100000000; i++ {
		intAttr[i%10000] = i
	}
	c := 0
	a := time.Now().UnixNano() / 1e6
	for i := 0; i < 100000000; i++ {
		c += intAttr[i%10000].(int)
	}
	log.Debug("cc:%v", time.Now().UnixNano()/1e6-a)
	intAttr1 := make(map[int]int64)
	for i := 0; i < 100000000; i++ {
		intAttr1[i%10000] = int64(i)
	}
	c1 := int64(0)
	a = time.Now().UnixNano() / 1e6
	for i := 0; i < 100000000; i++ {
		c1 += intAttr1[i%10000]
	}
	log.Debug("cc:%v", time.Now().UnixNano()/1e6-a)
}
