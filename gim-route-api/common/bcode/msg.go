/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-15 17:07:19
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-15 17:56:23
 */
package bcode

var msgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR:                       "fail",
	INVAILD_PARAMS:              "请求参数错误",
	USER_IS_EXSIT:               "用户已经存在",
	USERNAME_OR_PASSWORD_FAILED: "用户名或密码错误",
	ERROR_LOST_TOKEN:            "未传递token",
	ERROR_AUTH_CHECK_EXPIRED:    "token已经失效",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "token不合法",
}

func GetMsg(code int) string {
	if msg, ok := msgFlags[code]; !ok {
		return msgFlags[ERROR]
	} else {
		return msg
	}
}
