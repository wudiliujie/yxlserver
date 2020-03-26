package db

import (
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"yxlserver/services/I/eventcode"
)

var Module = new(DbModule)

type DbModule struct {
	*module.Skeleton
}

const DB_NUM = 10 //数据库线程数量
var dbMap = make(map[int64]*Db)

func (m *DbModule) OnInit() {
	m.Skeleton = module.NewSkeleton()
	m.Skeleton.Engine = m
	m.Name = "DbModule"
	Init()
	module.RegisterChanRpcEvent(eventcode.DB_LoadRoleData, m.ChanRPCServer, OnLoadRoleData)
}

func (m *DbModule) OnDestroy() {
	for k, v := range dbMap {
		v.Close()
		delete(dbMap, k)
	}
	waitClose.Wait()
	log.Debug("dbModule destroy end")
}
func Init() {
	for i := int64(0); i < DB_NUM; i++ {
		dbMap[i] = CreateDb(i)
	}
}
