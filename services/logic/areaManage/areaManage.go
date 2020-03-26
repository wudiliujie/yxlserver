package areaManage

import (
	"github.com/wudiliujie/common/eventhub"
	"yxlserver/services/I/eventcode"
	"yxlserver/services/consts"
	"yxlserver/services/tb/tbArea"
)

//小区管理

type AreaData struct {
	AreaId   int32 //区编号
	State    consts.ServerState
	AreaInfo *tbArea.ServerInfo
}

func (a *AreaData) CheckOpen() {
	if a.State == consts.ServerState_Run || a.State == consts.ServerState_Starting {
		return
	}
	//发出指令，加载区服数据
	//module.OnChanRpcEvent()

}
func (a *AreaData) CheckClose() {
	if a.State == consts.ServerState_Close {
		return
	}
	//保存数据,关闭
	eventhub.Publish(eventcode.Area_Close, a.AreaId)
}

var areaMap = make(map[int32]*AreaData)

func Init() {
	for _, v := range tbArea.GetAreaMap() {
		data := GetAreaData(v.ServerId)
		data.AreaInfo = v
		if v.State == 1 { //是开启状态
			data.CheckOpen()
		} else {
			data.CheckClose()
		}
	}
}

func GetAreaData(areaId int32) *AreaData {
	v, ok := areaMap[areaId]
	if !ok {
		v = &AreaData{}
		v.AreaId = areaId
		areaMap[areaId] = v
	}
	return v
}
func GetAreaState(areaId int32) consts.ServerState {
	v, ok := areaMap[areaId]
	if ok {
		return v.State
	}
	return consts.ServerState_Close
}
