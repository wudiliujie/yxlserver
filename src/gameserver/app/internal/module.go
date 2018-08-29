package internal

import (
	"gameserver/base"
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
}
func (m *AppModule) OnDestroy() {

}

func (m *AppModule) Update(diff int64) {

}
