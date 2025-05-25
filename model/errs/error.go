package errs

import "fmt"

const (
	CodeParamInvalid = 1001
	CodeServerError  = 1002
)

var (
	ParamInvalid = BizError{code: CodeParamInvalid, message: "参数错误"}
	ServerError  = BizError{code: CodeServerError, message: "服务器错误"}
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
	e.message = e.message + ":" + fmt.Sprintf(format, a...)
	return e
}
