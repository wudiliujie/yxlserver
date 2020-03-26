package db

import (
	"database/sql"
	"github.com/golang/protobuf/proto"
	"github.com/wudiliujie/common/convert"
	"github.com/wudiliujie/common/log"
	"github.com/wudiliujie/common/module"
	"github.com/wudiliujie/common/mysql"
	"github.com/wudiliujie/common/rtimer"
	"sync"
	"yxlserver/services/I/eventcode"
	"yxlserver/services/errmsg"
	"yxlserver/services/msg"
)

type Db struct {
	id int64
	*module.Skeleton
	closeSig chan bool
	timer    *rtimer.RTimer
}

var waitClose = sync.WaitGroup{}

//创建1个场景
func CreateDb(dbId int64) *Db {
	d := &Db{}
	d.id = dbId
	d.OnInit()
	d.Name = "db:" + convert.ToString(dbId)
	d.closeSig = make(chan bool, 1)
	d.Engine = d
	go func() {
		waitClose.Add(1)
		d.Run(d.closeSig)
		waitClose.Done()
		log.Debug("close db:%v", d.id)
	}()
	return d
}

func (d *Db) OnInit() {
	log.Debug("onInit:%v", d.id)
	d.Skeleton = module.NewSkeleton()
	d.RegisterChanRPC(eventcode.DB_LoadRoleData_DB, d.LoadRoleData)
}
func (d *Db) Close() {
	d.closeSig <- true
}

func (d *Db) LoadRoleData(args []interface{}) {
	ret := &msg.D2SLoadRoleCallback{}
	userId := args[0].(int64)
	defer func() {
		module.OnChanRpcEvent(eventcode.DB_LoadRoleData_Callback, ret)
	}()
	row, err := mysql.QueryRow("select * from user_info where user_id =?", userId)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error("LoadRoleData:%v", err)
			ret.Tag = errmsg.SYS_数据库错误
		} else {
			ret.Tag = errmsg.SYS_角色不存在
		}
		return
	}
	ret.UserId = userId
	ret.LoginAreaId = row.GetInt64("area_id")
	ret.Data = d.getRoleData(userId)
}
func (d *Db) getRoleData(userId int64) *msg.RoleDbData {
	row, err := mysql.QueryRow("select * from user_data where user_id =?", userId)
	ret := &msg.RoleDbData{}
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error("getRoleData:%v", err)
		} else {
			//插入一条数据
			_, err := mysql.Exec("insert into user_data (user_id,user_data)values(?,?)", userId, []byte{})
			if err != nil {
				log.Error("getRoleData 插入一条数据:%v", err)
			}
		}
	} else {
		data := row.GetBytes("user_data")
		err := proto.Unmarshal(data, ret)
		if err != nil {
			log.Error("getRoleData proto.Unmarshal %v", err)
		}
	}
	return ret
}
