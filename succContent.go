package boilerplate

import (
	"net/http"

	"github.com/woohaha/gin_boilerplate/errorCode"
)

type BaseResponse struct {
	ErrorCode int         `json:"error_code"`
	Msg       interface{} `json:"msg"`
}

func SuccResp(payload interface{}) (int, *BaseResponse) {
	return http.StatusOK, &BaseResponse{ErrorCode: errorCode.OK, Msg: payload}
}
