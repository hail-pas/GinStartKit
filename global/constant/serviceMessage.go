package constant

const (
	CodeSuccess = 0
	CodeError   = 1000

	CodeBadRequest   = 400
	CodeUnauthorized = 401
	CodeForbidden    = 403

	CodeErrorTokenParseFailed = 10001
	CodeErrorTokenExpire      = 10001
	CodeErrorSingleLoginOnly  = 10001
)

const (
	MessageSuccess = "success"
	MessageError   = "fail"

	MessageBadRequest   = "无效请求"
	MessageUnauthorized = "未授权"
	MessageForbidden    = "无权限"

	MessageTokenParseFailed    = "token解析失败"
	MessageTokenExpire         = "token过期"
	MessageSingleLoginOnly     = "只能单用户登录"
	MessageTokenGenerateFailed = "token生成失败"
	MessageRetrieveUserFailed  = "获取用户信息失败"
	UserWithUsernameExisted    = "该用户名已被使用"
	UserWithPhoneExisted       = "该手机号已被使用"

	MessagePasswordNull = "密码不能为空"
)
