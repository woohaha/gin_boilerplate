package errorCode

const (
	OK              = 0    // 成功
	ServerError     = 1000 // 系统错误
	NotFound        = 1001 // 401错误
	UnknownError    = 1002 // 未知错误
	ParameterError  = 1003 // 参数错误
	PermissionError = 1004 // 認證错误
	ForbiddenError  = 1005 // 禁止訪問
)
