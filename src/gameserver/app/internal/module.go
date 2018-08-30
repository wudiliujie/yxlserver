package internal

import (
	"gameserver/base"
	"gameserver/common/playerManage"
	"leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type AppModule struct {
	*module.Skeleton
}

func (m *AppModule) OnInit() {
	m.Skeleton = skeleton
	m.Skeleton.Engine = m
	base.APP = m.Skeleton
}
func (m *AppModule) OnDestroy() {

}

func (m *AppModule) Update(diff int64) {
	playerManage.Update(diff)
}
