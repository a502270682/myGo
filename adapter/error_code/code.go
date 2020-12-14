package error_code

const (
	CODE_UNDEFINED    = -1
	CODE_OK           = 0
	CODE_SYSTEM_ERROR = 1
	CODE_PARAM_WRONG  = 2
)

var (
	errCodeMsg = map[int]string{
		CODE_UNDEFINED:    "未知错误",
		CODE_OK:           "成功",
		CODE_SYSTEM_ERROR: "系统错误",
		CODE_PARAM_WRONG:  "参数错误",
	}
)

func ErrCodeMessage(code int) string {
	return errCodeMsg[code]
}
