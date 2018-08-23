package rtime

import "time"

//引擎 时间工具

var (
	startTime  time.Time     //启动时间
	runTime    int64         //运行时间
	repairTime int64     = 0 //修复时间
)

func init() {
	startTime = time.Now()
}
func Now() time.Time {
	if repairTime != 0 {
		return time.Now().Add(time.Microsecond * time.Duration(repairTime))
	} else {
		return time.Now()
	}
}

//设置修复时间
func SetRepairTime(v int64) {
	repairTime = v
}
