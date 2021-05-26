package error_code

const (
	CodeUndefined   = -1
	CodeOk          = 0
	CodeSystemError = 1
	CodeParamWrong  = 2
)

var (
	errCodeMsg = map[int]string{
		CodeUndefined:   "未知错误",
		CodeOk:          "成功",
		CodeSystemError: "系统错误",
		CodeParamWrong:  "参数错误",
	}
)

func ErrCodeMessage(code int) string {
	return errCodeMsg[code]
}
