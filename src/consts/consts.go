package consts
type ROOM_TYPE int32
const (
	ROOM_TYPE_NORMAL    ROOM_TYPE = 0  //正常房间
	ROOM_TYPE_MATCH   	ROOM_TYPE = 1  //比赛房间
)
const (
	ROOM_MAX_NUM =4	//房间最大人数
)


const(
	Center_Rpc_GetBestRoomId="GetBestRoomId" //获取一个房间编号
	Center_Rpc_GetModuleByRoomId="GetModuleByRoomId" //根据房间获取module
	Center_Rpc_GetNewRoomId="GetNewRoomId" //获取新的房间编号
	Center_Rpc_OnCreateRoom="OnCreateRoom" //当创建房间
	Center_Rpc_OnRoomNumChange="OnRoomNumChange" //当房间人数发生变化

	Center_Rpc_OnPlayerChangeModule="OnPlayerChangeModule" //player module发生变化

	Room_Rpc_LoginModule="LoginModule" //进入module
	Room_Rpc_EnterModule="EnterModule" //进入module
	Room_Rpc_EnterModuleRoom="EnterModuleRoom" //进入房间
)