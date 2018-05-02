package e

const (
	SUCCESS = 000000

	USER_INVALID_PARAMS = 100001
	USER_CAPTCHA_ERROR = 100002
	USER_AUTH_CHECK_TOKEN_FAIL = 100003
	USER_AUTH_CHECK_TOKEN_TIMEOUT = 100004
	USER_VCODE_ERROR = 100005
	USER_INTER_ERROR = 100006
	USER_LOGIN_ERROR = 100007


	ERROR_UNKOWN = 999999

)

var MsgFlag = map[int]string {
	SUCCESS: "ok",

	USER_INVALID_PARAMS: "请求参数错误",
	USER_CAPTCHA_ERROR: "图片验证码失效",
	USER_AUTH_CHECK_TOKEN_FAIL: "token错误",
	USER_AUTH_CHECK_TOKEN_TIMEOUT: "token失效",
	USER_VCODE_ERROR: "验证码错误",
	USER_INTER_ERROR: "内部错误",
	USER_LOGIN_ERROR: "用户不存在或密码错误",


	ERROR_UNKOWN: "未知错误",
}

func GetErrorMessage(code int) string{
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[ERROR_UNKOWN]
}
