package boilerplate

import (
	"fmt"
	"net/http"

	"github.com/woohaha/gin_boilerplate/errorCode"
)

// ServerError 內部错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, errorCode.ServerError, http.StatusText(http.StatusInternalServerError))
}

// NotFound 資源不存在错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, errorCode.NotFound, http.StatusText(http.StatusNotFound))
}

// PermissionError 認證錯誤
func PermissionError() *APIException {
	return newAPIException(http.StatusUnauthorized, errorCode.PermissionError, http.StatusText(http.StatusUnauthorized))
}

// UnknownError 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, errorCode.UnknownError, message)
}

// ParameterError 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, errorCode.ParameterError, message)
}

type APIException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

func (e *APIException) Error() string {
	return e.Msg
}

func (e *APIException) setRequestURI(r *http.Request) {
	e.Request = fmt.Sprintf("%s %s", r.Method, r.URL.String())
}

func newAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}
