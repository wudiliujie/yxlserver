package rooms

import (
	"gameserver/common"
	"consts"
)



var (
	RoomSN int //房间序号
	roomMap map[int]*RoomData
)
type RoomData struct {
	Id int32
	RoomType consts.ROOM_TYPE
	Module common.RoomModule
	Num  int 						//房间人数
}

func GetNewRoomId() int{
	RoomSN++
	return  RoomSN
}
func GetBestRoomId(args []interface{}) (interface{},error){
	t:=args[0].(consts.ROOM_TYPE)
	var id int32 =0
	for _,v :=range roomMap{
		if v.RoomType== t{
			if v.Num<consts.ROOM_MAX_NUM{
				id=v.Id
			}
		}
	}
	return  id,nil
}
func GetModuleByRoomId(args[] interface{})(interface{},error){
	return  nil,nil
}


