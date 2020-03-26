package tb

import (
	"encoding/json"
)

type Cfg_Area struct {
	AreaId    int32
	AreaName  string
	StartDate int32
	State     int32
}
type CfgData struct {
	Ver   string
	Items *CfgDataAll
}
type CfgDataAll struct {
	Cfg_Area []*Cfg_Area
}

var Data *CfgDataAll
var Ver string

func InitCfgData(data []byte) error {
	var errcode error
	var allData CfgData
	errcode = json.Unmarshal(data, &allData)
	if errcode != nil {
		return errcode
	}
	Ver = allData.Ver
	Data = allData.Items
	return nil
}
