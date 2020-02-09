package eventcode

const (
	Net_AgentInit        = "Net_AgentInit"
	Net_AgentDestroy     = "Net_AgentDestroy"
	Net_ReceiveMsg       = "Net_ReceiveMsg"
	Net_ClientInit       = "Net_ClientInit" //客户端初始化成功
	Net_ClientDestroy    = "Net_ClientDestroy"
	Net_ClientReceiveMsg = "Net_ClientReceiveMsg"
	Login                = "Login"

	DB_LoadRoleData          = "DB_LoadRoleData"
	DB_SaveRoleData          = "DB_SaveRoleData"
	DB_LoadRoleData_DB       = "DB_LoadRoleData_DB"
	DB_LoadRoleData_CallBack = "DB_LoadRoleData_CallBack"
	DB_SaveRoleData_DB       = "DB_SaveRoleData_DB"
	DB_GetUserMail           = "DB_GetUserMail"          //获取用户邮件
	DB_GetUserMail_DB        = "DB_GetUserMail_DB"       //获取用户邮件
	DB_GetUserMail_Callback  = "DB_GetUserMail_Callback" //获取用户邮件
	DB_SaveUserMail          = "DB_SaveUserMail"         //用户保存邮件
	DB_SaveUserMail_DB       = "DB_SaveUserMail_DB"      //用户保存邮件

	DB_GetToken         = "DB_GetToken"
	DB_GetToken_DB      = "DB_GetToken_DB"
	DB_GetTokenCallBack = "DB_GetTokenCallBack"

	DB_SaveGang    = "DB_SaveGang"    //帮会落库
	DB_SaveGang_DB = "DB_SaveGang_DB" //帮会落库
	DB_RedPkg      = "DB_RedPkg"      //红包落库

	DB_SendShopSave    = "DB_SendShopSave"    //寄售保存
	DB_SendShopLogSave = "DB_SendShopLogSave" //寄售日志保存
	DB_SendShopLogDel  = "DB_SendShopLogDel"  //寄售日志删除

	DB_GetUserInfo          = "DB_GetUserInfo"          //赠送查询用户信息
	DB_GetUserInfoDB        = "DB_GetUserInfoDB"        //赠送查询用户信息
	DB_GetUserInfo_Callback = "DB_GetUserInfo_Callback" //寄售服务器查询用户信息回调
	//赠送
	DB_SendItem         = "DB_SendItem"         //赠送
	DB_SendItemDB       = "DB_SendItemDB"       //赠送
	DB_SendItemCallback = "DB_SendItemCallback" //赠送回调，如果失败，则退回用户道具，修改赠送数量

	DB_SaveServerData    = "DB_SaveServerData"    //保存服务器数据
	DB_SaveServerData_DB = "DB_SaveServerData_DB" //保存服务器数据

	DB_SaveCrossServerData    = "DB_SaveCrossServerData"    //保存跨服服务器数据
	DB_SaveCrossServerData_DB = "DB_SaveCrossServerData_DB" //保存跨服服务器数据

	DB_SaveAndGetStock          = "DB_SaveAndGetStock"          //保存库存，并获取新的库存
	DB_SaveAndGetStock_DB       = "DB_SaveAndGetStock_DB"       //保存库存，并获取新的库存
	DB_SaveAndGetStock_Callback = "DB_SaveAndGetStock_Callback" //保存库存，回调

	DB_SaveFriendData              = "DB_SaveFriendData"              //保存好友数据
	DB_SaveFriendData_DB           = "DB_SaveFriendData_DB"           //保存好友数据
	DB_SaveMarryData               = "DB_SaveMarryData"               //保存姻缘数据
	DB_SaveMarryData_DB            = "DB_SaveMarryData_DB"            //保存姻缘数据
	DB_SaveMasterData              = "DB_SaveMasterData"              //保存师徒数据
	DB_SaveMasterData_DB           = "DB_SaveMasterData_DB"           //保存师徒数据
	DB_SaveGodCluePlayer           = "DB_SaveGodCluePlayer"           //保存仙道会用户数据
	DB_SaveGodClubPlayer_DB        = "DB_SaveGodClubPlayer_DB"        //db线程保存
	DB_GodClubCalculation          = "DB_GodClubCalculation"          //计算战斗
	DB_GodClubCalculation_DB       = "DB_GodClubCalculation_DB"       //计算战斗
	DB_GodClubCalculation_CallBack = "DB_GodClubCalculation_CallBack" //计算战斗，返回每个区的报名人数

	DB_SaveGodClubBattlefield               = "DB_SaveGodClubBattlefield"               //保存仙道会用户数据
	DB_SaveGodClubBattlefield_DB            = "DB_SaveGodClubBattlefield_DB"            //db线程保存
	DB_GetGodClubBattlefieldPlayer          = "DB_GetGodClubBattlefieldPlayer"          //db线程保存
	DB_GetGodClubBattlefieldPlayer_DB       = "DB_GetGodClubBattlefieldPlayer_DB"       //db线程保存
	DB_GetGodClubBattlefieldPlayer_CallBack = "DB_GetGodClubBattlefieldPlayer_CallBack" //db线程保存
	DB_SaveGodClubFightReport               = "DB_SaveGodClubFightReport"               //保存战报
	DB_SaveGodClubFightReport_DB            = "DB_SaveGodClubFightReport_DB"            //保存战报子线程

	DB_DelGodClubFightResult     = "DB_DelGodClubFightResult"     //删除结果
	DB_DelGodClubFightResult_DB  = "DB_DelGodClubFightResult_DB"  //删除结果
	DB_SaveGodClubFightResult    = "DB_SaveGodClubFightResult"    //保存结果
	DB_SaveGodClubFightResult_DB = "DB_SaveGodClubFightResult_DB" //保存结果

	DB_LoadGodClubFightResult          = "DB_LoadGodClubFightResult"          //保存结果
	DB_LoadGodClubFightResult_DB       = "DB_LoadGodClubFightResult_DB"       //保存结果
	DB_LoadGodClubFightResult_CallBack = "DB_LoadGodClubFightResult_CallBack" //保存结果

	DB_SaveNineData    = "DB_SaveNineData"    //保存九重天数据
	DB_SaveNineData_DB = "DB_SaveNineData_DB" //保存九重天数据

	DB_GetNineData          = "DB_GetNineData"
	DB_GetNineData_DB       = "DB_GetNineData_DB"       //获取九重天数据
	DB_GetNineData_Callback = "DB_GetNineData_Callback" //获取九重天数据callback

	Http_HttpReq         = "Http_HttpReq"
	Http_HttpGet         = "Http_HttpGet"
	Http_HttpPostForm    = "Http_HttpPostForm"
	Http_UpdateLoginTime = "Http_UpdateLoginTime" //更新用户登录时间

	Cross_RpcSendPack = "Cross_RpcSendPack"
	Cross_SendPack    = "Cross_SendPack"

	Cross_SendServerPack = "Cross_SendServerPack"

	DB_Activity  = "DB_Activity"  //活动数据落地
	DB_New_Rank  = "DB_New_Rank"  //新版排行活动数据落地
	DB_DailyTask = "DB_DailyTask" //日常数据落地

	App_GetOnlineUser = "App_GetOnlineUser"
	App_GetUserInfo   = "App_GetUserInfo"
	App_Pay           = "App_Pay" //充值
	App_CloseServer   = "App_CloseServer"
	App_SendMail      = "App_SendMail"    //发送邮件
	App_SendItem      = "App_SendItem"    //增送
	App_GM            = "App_GM"          //执行gm
	App_RestoreData   = "App_RestoreData" //执行gm
	App_Chat          = "App_Chat"        //执行聊天
)
