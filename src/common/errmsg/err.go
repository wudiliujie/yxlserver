package errmsg
const (
	SYS_ASYNC_WAIT=-1	//异步等待
	SYS_SUCCESS=0	//成功
	SYS_LOGIN_NOACCOUNT=1	//帐号错误
	SYS_LOGIN_PASSWORD=2	//密码错误
	SYS_LOGIN_NO_MODULE=3	//没有可用的module
	SYS_LOGIN_ENTER_MODULE_ERROR=4	//进入模块错误
	SYS_LOGIN_ALREAD_ONLINE=5	//重复登录
	SYS_ROOM_SYSTEM_ERROR=101	//系统错误
)
