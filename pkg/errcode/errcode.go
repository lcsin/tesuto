package errcode

import "errors"

//go:generate stringer -type=ErrCode -linecomment
type ErrCode int64

const (
	OK           ErrCode = 0  // ok
	UNKnownError ErrCode = -1 // unknown error
	InvalidParam ErrCode = -2 // 参数无效
	UnAuthorized ErrCode = -3 // 未授权

	UserNotFound           ErrCode = -1001 // 用户不存在或密码错误
	EmailAlreadyRegistered ErrCode = -1002 // 邮箱已注册
	PasswordInconsistency  ErrCode = -1003 // 密码不一致
	EmailIsEmpty           ErrCode = -1004 // 邮箱不能为空
	UserNotExists          ErrCode = -1005 // 用户不存在
)

var (
	ErrInternalServer = errors.New("系统错误")
	ErrUnAuthorized   = errors.New("未授权")
	ErrInvalidParams  = errors.New("参数错误")
)
