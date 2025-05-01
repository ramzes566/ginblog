package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NOT_RIGHT   = 1008
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "Token不存在，请重新登录",
	ERROR_TOKEN_RUNTIME:    "Token已过期，请重新登录",
	ERROR_TOKEN_WRONG:      "Token错误，请重新登录",
	ERROR_TOKEN_TYPE_WRONG: "Token格式错误，请重新登录",
	ERROR_USER_NOT_RIGHT:   "用户权限不足",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
