// Code generated by "stringer -type=ErrCode -linecomment"; DO NOT EDIT.

package errcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OK-0]
	_ = x[UNKnownError - -1]
	_ = x[InvalidParam - -2]
	_ = x[UnAuthorized - -3]
	_ = x[UserNotFound - -1001]
	_ = x[EmailAlreadyRegistered - -1002]
	_ = x[PasswordInconsistency - -1003]
	_ = x[EmailIsEmpty - -1004]
	_ = x[UserNotExists - -1005]
}

const (
	_ErrCode_name_0 = "用户不存在邮箱不能为空密码不一致邮箱已注册用户不存在或密码错误"
	_ErrCode_name_1 = "未授权参数无效unknown errorok"
)

var (
	_ErrCode_index_0 = [...]uint8{0, 15, 33, 48, 63, 93}
	_ErrCode_index_1 = [...]uint8{0, 9, 21, 34, 36}
)

func (i ErrCode) String() string {
	switch {
	case -1005 <= i && i <= -1001:
		i -= -1005
		return _ErrCode_name_0[_ErrCode_index_0[i]:_ErrCode_index_0[i+1]]
	case -3 <= i && i <= 0:
		i -= -3
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
