package I

import "yxlserver/services/consts"

type IUnit interface {
	GetId() int64
	Init()
	GetIntValue(attrType consts.AttrType) int64
	SetIntValue(attrType consts.AttrType, v int64)
	AddIntValue(attrType consts.AttrType, v int64)
	GetStrValue(attrType consts.AttrType) string
	SetStrValue(attrType consts.AttrType, v string)
	GetInt32Value(attrType consts.AttrType) int32
}
