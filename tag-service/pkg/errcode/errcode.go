package errcode

import (
	"fmt"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var _codes = map[int]string{} //命名与grpc/codes 冲突故而加_

func NewError(code int, msg string) *Error {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在，请更换一个", code))
	}
	_codes[code] = msg

	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d,错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details []string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

