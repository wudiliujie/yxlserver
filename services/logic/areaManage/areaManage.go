package areaManage

import "yxlserver/services/consts"

//小区管理

type AreaInfo struct {
	AreaId    int32  //区编号
	AreaName  string //区名字
	State     consts.AreaState
	StartDate int32 //开区时间
}
