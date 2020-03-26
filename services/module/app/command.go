package app

import (
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"yxlserver/services/I"
	"yxlserver/services/logic/playerManage"
)

func initCommand(s *module.Skeleton) {
	s.RegisterCommand("load", "login account: input accountName, passward", Debug)
}
func Debug(args []interface{}) (ret interface{}, err error) {
	playerManage.SignPlayerAsyncFunc(3, func(targetPlayer I.IPlayer, args ...interface{}) {
		log.Debug("callback:")
	})
	return "success", nil
}
