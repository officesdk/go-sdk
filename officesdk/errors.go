package officesdk

import (
	"fmt"
)

const (
	OK     int = 0
	Failed int = 1
)

type Error struct {
	httpStatusCode int
	code           int
	message        string
}

// Code 获取错误码
func (err *Error) Code() int {
	return err.code
}

// HttpStatusCode 获取错误码对应的HTTP状态码
func (err *Error) HttpStatusCode() int {
	return err.httpStatusCode
}

// Message 获取错误描述
func (err *Error) Message() string {
	return err.message
}

// WithMessage 设置错误描述
func (err *Error) WithMessage(msg string) *Error {
	clone := *err
	err.message = msg
	return &clone
}

// Error 获取错误描述
func (err *Error) Error() string {
	return fmt.Sprintf("code:%d message:%s", err.code, err.message)
}

// NewError 创建枚举错误
func NewError(httpStatusCode int) *Error {
	return &Error{httpStatusCode: httpStatusCode, code: Failed}
}

// NewCustomError 创建自定义错误
func NewCustomError(httpStatusCode, code int, message string) *Error {
	return &Error{httpStatusCode: httpStatusCode, code: code, message: message}
}
