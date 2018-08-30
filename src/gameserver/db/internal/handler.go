package internal

import (
	"common/proto"
	"gameserver/common/mysql"
	"gameserver/events"
	p1 "github.com/golang/protobuf/proto"
	"leaf/log"
	"time"
)

func init() {
	ChanRPC.Register(events.DB_LoadRoleData, OnLoadRoleData)
}
func OnLoadRoleData(args []interface{}) (interface{}, error) {
	log.Debug("load user")

	rows, err := mysql.Query("select * from user_info where userid =?", 1)
	if err != nil {
		log.Error("%v>>", err)
	}
	roleinfo := &proto.RoleDbInfo{}
	p1.UnmarshalMerge(rows[0].GetBytes("UserData"), roleinfo)
	if len(rows) > 0 {
		log.Debug("%v", rows[0].GetInt32("UserId"))
		log.Debug("%v", rows[0].GetBytes("UserData"))
	}
	log.Debug("%v", roleinfo)
	time.Sleep(time.Second * 5)
	return roleinfo, nil
}
