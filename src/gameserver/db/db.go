package db

import(
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"leaf/log"
)
func GetDb() *sql.DB{
	db,err:= sql.Open("mysql","root:111111@tcp(127.0.0.1:3306)/yxl_game?charset=utf8&parseTime=true")
	if(err!=nil){
		log.Error("数据库错误%v",err)
	}
	return db;
}
func ReadRoleInfo(roleId int32,data *[]byte)error {
	db := GetDb();
	 row:= db.QueryRow("select role_data from user_account where roleid=?",roleId)
	 if(row!=nil){
	 	row.Scan(&data);
	 }
	return  nil;
}
func SaveRoleInfo(roleId int32,data *[]byte){
	db:=GetDb();
	_, err:= db.Exec("insert into user_account (roleid,role_data)values(?,?)",roleId,data);
	if(err!=nil){
		log.Error("保存玩家错误%v",err)
	}

}
