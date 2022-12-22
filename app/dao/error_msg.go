package dao

import "errors"

var (
	UserExists        = errors.New("用户已存在")
	UserNotExists     = errors.New("用户不存在")
	PasswordError     = errors.New("密码错误")
	TokenExpiredError = errors.New("token 已过期")
	DataNotExists     = errors.New("数据为空")
	VoteTimeExpired   = errors.New("投票时间已过")
	VoteRepeated      = errors.New("不能重复投票")
)
