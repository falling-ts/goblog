package error

import "fmt"

type Error struct {
	Code int
	Msg  string
}

// NewError 创建错误实例
func NewError() *Error {
	return &Error{}
}

// Error 错误格式化
func (err *Error) Error() string {
	return fmt.Sprintf("[%d]Error Message: %s", err.Code, err.Msg)
}
