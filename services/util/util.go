package util

import (
	"fmt"
	"github.com/wudiliujie/common/security"
	"reflect"
)

//获取登录签名
func GetLoginSign(userId int64, loginTime int64) string {
	str := fmt.Sprintf("%v%v%v", userId, loginTime, "111111")
	return security.Md5(str)
}
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsNil()
}
