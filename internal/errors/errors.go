package errors

import (
	"fmt"
)

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error(%d): %s", e.Code, e.Msg)
}

// New 创建一个errcode
func New(code int, msg string) error {
	return &Error{Code: code, Msg: msg}
}
