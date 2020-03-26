package tbArea

import (
	"github.com/wudiliujie/common/log"
	"yxlserver/services/tb"
)

type ServerInfo struct {
	ServerId   int32  //区编号
	ServerName string //区名字
	State      int32
	StartDate  int32 //开区时间
}

var serverMap = make(map[int32]*ServerInfo)

func Init() {
	log.Release("初始化区信息")
	_areaMap := make(map[int32]*ServerInfo)
	for _, v := range tb.Data.Cfg_Area {
		info := &ServerInfo{}
		info.ServerId = v.AreaId
		info.ServerName = v.AreaName
		info.State = v.State
		info.StartDate = v.StartDate
		_areaMap[info.ServerId] = info
	}
	serverMap = _areaMap
}
func GetAreaMap() map[int32]*ServerInfo {
	return serverMap
}
