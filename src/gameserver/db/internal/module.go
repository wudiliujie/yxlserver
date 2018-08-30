package internal

import (
	"gameserver/base"
	"leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type DbModule struct {
	*module.Skeleton
}

func (m *DbModule) OnInit() {
	m.Skeleton = skeleton
	m.Skeleton.Engine = m
	base.DB = m.Skeleton
}
func (m *DbModule) OnDestroy() {

}
