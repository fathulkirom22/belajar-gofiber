package blueprint

import "fmt"

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("status %d: %v", err.Code, err.Msg)
}

func CreateError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
