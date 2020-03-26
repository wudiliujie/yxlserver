package errmsg

type ErrCode int32

const (
	SYS_成功        = 0 //成功
	SYS_签名错误      = 1 //签名错误
	SYS_读取数据错误    = 2 //读取数据错误
	SYS_服务器维护中    = 3 //服务器维护中
	SYS_帐号在其他地方登录 = 4 //帐号在其他地方登录
	SYS_角色不存在     = 5 //角色不存在
	SYS_数据库错误     = 6 //数据库错误
)
