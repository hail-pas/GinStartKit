package constant

const (
	CodeSuccess = 0
	CodeError   = 1000

	CodeInvalidRequest = 400 + iota
	CodeUnauthorized

	CodeErrorTokenParseFailed = 10001 + iota
	CodeErrorTokenExpire
	CodeErrorSingleLoginOnly
)

const (
	MessageSuccess = "success"
	MessageError   = "fail"

	MessageInvalidRequest   = "无效请求"
	MessageCodeUnauthorized = "未授权"

	MessageTokenParseFailed    = "Token解析失败"
	MessageTokenExpire         = "Token过期"
	MessageSingleLoginOnly     = "只能单用户登录"
	MessageTokenGenerateFailed = "Token生成失败"
	MessageRetrieveUserFailed  = "获取用户信息失败"

	MessagePasswordNull = "密码不能为空"
)
