package controllers

const (
	CodeOk            = 200 // 正常
	CodeInternalError = 500 // 系统内部错误

	CodeBadRequest      = 40000 // 非法请求
	CodeUnauthorized    = 40100 // 请先登录
	CodeForbidden       = 40300 // 禁止访问
	CodeDataNotExists   = 40400 // 数据不存在
	CodeEmptyParam      = 41600 // 参数为空
	CodeUnexpectedParam = 41700 // 参数错误
	CodeInvalidToken    = 41800 // 无效的 token
	CodeTokenExpired    = 41900 // token 已过期
	CodeVoteTimeExpired = 42000 // 投票时间已过
	CodeVoteRepeated    = 42100 // 不能重复投票
	CodeTooManyAttempts = 42200 // 操作频繁

	CodeServiceExpired  = 50401 // 服务已过期
	CodeServiceRepeated = 50402 // 重复投票
	CodeUserExists      = 51201 // 用户已存在
	CodeUserNotExists   = 51202 // 用户不存在
	CodePasswordError   = 51300 // 密码错误
)

var codeMsgMap = map[int]string{
	CodeOk:              "成功",
	CodeUnauthorized:    "请先登录",
	CodeUnexpectedParam: "请求参数错误",
	CodeInvalidToken:    "无效的 token",
	CodeTokenExpired:    "token 已过期",
	CodeVoteTimeExpired: "投票时间已过",
	CodeVoteRepeated:    "不能重复投票",
	CodeTooManyAttempts: "操作频繁",
	CodeDataNotExists:   "数据不存在",
	CodeInternalError:   "服务繁忙",
	CodeUserExists:      "用户已存在",
}

func MsgTender(code int) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeInternalError]
	}
	return msg
}
