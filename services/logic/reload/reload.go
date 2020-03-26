package reload

import (
	"github.com/wudiliujie/common/log"
	"io/ioutil"
	"yxlserver/services/tb"
	"yxlserver/services/tb/tbArea"
)

func Reload() {
	data, errcode := ioutil.ReadFile("conf/cfg_data.json")
	if errcode != nil {
		log.Error("bReload  %v", errcode)
		return
	}
	tb.InitCfgData(data)
	tbArea.Init()
}
