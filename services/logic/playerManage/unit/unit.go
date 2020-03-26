package unit

import (
	"yxlserver/services/I"
	"yxlserver/services/consts"
)

//属性管理
type Attr struct {
	Value1 int64
	Value2 string
	Change bool
}

//基础单位
type Unit struct {
	Id       int64
	Self     I.IUnit
	Attr     map[consts.AttrType]*Attr
	UnitType consts.GradeType
}

func (u *Unit) Init() {
	u.Attr = make(map[consts.AttrType]*Attr)
}
func (u *Unit) GetAttr(attrType consts.AttrType) *Attr {
	a, ok := u.Attr[attrType]
	if !ok {
		a = &Attr{}
		u.Attr[attrType] = a
	}
	return a
}

func (u *Unit) GetId() int64 {
	return u.Id
}
func (u *Unit) GetIntValue(attrType consts.AttrType) int64 {
	return u.GetAttr(attrType).Value1
}
func (u *Unit) GetInt32Value(attrType consts.AttrType) int32 {
	return int32(u.GetAttr(attrType).Value1)
}
func (u *Unit) SetIntValue(attrType consts.AttrType, v int64) {
	a := u.GetAttr(attrType)
	a.Value1 = v
	a.Change = true
}
func (u *Unit) AddIntValue(attrType consts.AttrType, v int64) {
	a := u.GetAttr(attrType)
	a.Value1 += v
	a.Change = true
}
func (u *Unit) GetStrValue(attrType consts.AttrType) string {
	return u.GetAttr(attrType).Value2
}
func (u *Unit) SetStrValue(attrType consts.AttrType, v string) {
	a := u.GetAttr(attrType)
	a.Value2 = v
	a.Change = true
}
