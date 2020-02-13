package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"yxlserver/services/conf"
	"yxlserver/services/consts"
	"yxlserver/services/module/app"
	"yxlserver/services/module/gate"
	"yxlserver/services/msg"
)

func main() {
	Test()
	conf.Server.ServerId = 1
	conf.Server.ServerType = consts.ServerType_Gate
	conf.Server.ServerName = "网关"
	modules := []module.Module{app.CreateGateAppModule(), gate.Module}
	/*	if conf.Server.DebugAddr != "" {
		log.Release("启动debug端口:%v", conf.Server.DebugAddr)
		go http.ListenAndServe(conf.Server.DebugAddr, nil)
	}*/
	module.Run(modules...)
}
func Test() {
	a := &msg.S2CRoleInfo{}
	a.UserId = 100
	a.A = append(a.A, &msg.IntAttr{K: 100, V: 100})
	a.A = append(a.A, &msg.IntAttr{K: 101, V: 101})
	a.A = append(a.A, &msg.IntAttr{K: 102, V: 102})
	data, err := proto.Marshal(a)
	if err != nil {
		log.Debug("aaa:%v", err)
	}
	b := &msg.S2CRoleInfo{}
	err = proto.Unmarshal(data, b)
	if err != nil {
		log.Debug("aaa:%v", err)
	}

}
