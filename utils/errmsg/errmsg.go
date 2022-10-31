package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块错误 1000...
	ERROR_USERNAME_USED   = 1001
	ERROR_PASSWORD_WRONG  = 1002
	ERROR_USER_NOT_EXIST  = 1003
	ERROR_TOKEN_NOT_EXIST = 1004
	ERROR_TOKEN_RUNTIME   = 1005
	ERROR_TOKEN_WRONG     = 1006
	ERROR_TOKEN_FORMAT    = 1007

	// 标签模块错误 2000...
	ERROR_TAGNAME_USED = 3001

	// 文章模块错误 3000...

)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "FAIL",
	ERROR_USERNAME_USED:   "用户名已存在",
	ERROR_PASSWORD_WRONG:  "密码错误",
	ERROR_USER_NOT_EXIST:  "用户不存在",
	ERROR_TOKEN_NOT_EXIST: "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:   "TOKEN已过期",
	ERROR_TOKEN_WRONG:     "TOKEN不正确",
	ERROR_TOKEN_FORMAT:    "TOKEN格式错误",
	ERROR_TAGNAME_USED:    "标签名已存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
