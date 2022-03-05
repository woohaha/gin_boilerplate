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

// Forbidden 禁止訪問
func Forbidden(message string) *APIException {
	if message == "" {
		message = http.StatusText(http.StatusForbidden)
	}
	return newAPIException(http.StatusForbidden, errorCode.ForbiddenError, message)
}

// PermissionError 認證錯誤
func PermissionError(message string) *APIException {
	if message == "" {
		message = http.StatusText(http.StatusUnauthorized)
	}
	return newAPIException(http.StatusUnauthorized, errorCode.PermissionError, message)
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
	BaseResponse
	Code    int    `json:"-"`
	Request string `json:"request"`
}

func (e *APIException) Error() string {
	return e.Msg.(string)
}

func (e *APIException) setRequestURI(r *http.Request) {
	e.Request = fmt.Sprintf("%s %s", r.Method, r.URL.String())
}

func newAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code: code,
		BaseResponse: BaseResponse{
			ErrorCode: errorCode,
			Msg:       msg,
		},
	}
}
