package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"leaf/log"
	"sync/atomic"
)
var (
	DBTestCount=int32(0)
	db *sql.DB
)
func Init()  {
	var err error
	db, err = sql.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/yxl_game?charset=utf8&parseTime=true")
	if err != nil {
		log.Error("数据库错误%v", err)
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}
func ReadRoleInfo(roleId int32, data *[]byte) error {
	row := db.QueryRow("select role_data from user_account where roleid=?", roleId)
	if row != nil {
		row.Scan(data)
	}
	return nil
}
func SaveRoleInfo(roleId int32, data *[]byte) {

	_, err := db.Exec("update user_account set role_data=? where roleid=?", data, roleId)
	if err != nil {
		log.Error("保存玩家错误%v", err)
		atomic.AddInt32(&DBTestCount,1)
		log.Debug("saveerr:%v>>>%v",atomic.LoadInt32(&DBTestCount))
	}
}
func CreateRoleInfo(roleId int32, data *[]byte) {
	_, err := db.Exec("insert into user_account (roleid,role_data)values(?,?)", roleId, data)
	if err != nil {
		log.Error("创建玩家错误%v", err)
	}
}

func GetAccountId(username string, userpass string) int32 {
	row := db.QueryRow("select id from user_login where username=?", username)

	var id int32
	row.Scan(&id)
	if id != 0 {
		return id
	}
	ret, err := db.Exec("insert into user_login (username,userpass)values(?,?)", username, userpass)
	if err != nil {
		log.Error("创建帐号错误:%v", err)
		return 0
	}
	id1, err := ret.LastInsertId()
	if err != nil {
		log.Error("创建帐号错误:%v", err)
		return 0
	}
	return int32(id1)
}
