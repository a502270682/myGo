package error_code

import "strings"

const autoPrefix = "ReplyError:"

type ReplyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ReplyError) IsAutoMsg() bool {
	return strings.HasPrefix(e.Message, autoPrefix) || e.Message == ""
}

func Error(code int, msg string) *ReplyError {
	return &ReplyError{
		Code:    code,
		Message: msg,
	}
}
