package db

import (
	"github.com/wudiliujie/common/log"
	"yxlserver/services/I/eventcode"
)

func OnLoadRoleData(args []interface{}) {
	userId := args[0].(int64)
	id := userId % DB_NUM
	db, ok := dbMap[id]
	if ok {
		db.ChanRPCServer.Go(eventcode.DB_LoadRoleData_DB, userId)
	} else {
		log.Error("db 不存在：%v >>%v", id, userId)
	}
}
