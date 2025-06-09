package model

import "fmt"

const (
	CodeParamInvalid     = 1001
	CodeServerError      = 1002
	CodeUserAlreadyExist = 1003
	CodeUserNotExist     = 1004
	CodeUserNotLogin     = 1005
	CodeTokenInvalid     = 1006
)

var (
	ParamInvalid     = &BizError{code: CodeParamInvalid, message: "参数错误"}
	ServerError      = &BizError{code: CodeServerError, message: "服务器错误"}
	UserAlreadyExist = &BizError{code: CodeUserAlreadyExist, message: "邮箱已被注册"}
	UserNotExist     = &BizError{code: CodeUserNotExist, message: "邮箱未注册"}
	UserNotLogin     = &BizError{code: CodeUserNotLogin, message: "用户未登录"}
	TokenInvalid     = &BizError{code: CodeTokenInvalid, message: "token无效"}
)

type BizError struct {
	code    int
	message string
}

func (e *BizError) Code() int {
	return 0
}
func (e *BizError) Message() string {
	return e.message
}
func (e *BizError) Error() string {
	return e.message
}
func (e *BizError) WithMessage(format string, a ...any) *BizError {
	return &BizError{
		code:    e.code,
		message: e.message + ":" + fmt.Sprintf(format, a...),
	}
}
