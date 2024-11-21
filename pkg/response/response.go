package response

import (
	"github.com/lcsin/tesuto/pkg/errcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Response 响应体封装
type Response struct {
	Code    errcode.ErrCode `json:"code"`
	Message string          `json:"message"`
	Data    any             `json:"data"`
}

// OK 请求成功
func OK(data any) {
	//c.JSON(http.StatusOK, Response{
	//	Code:    0,
	//	Message: "ok",
	//	Data:    data,
	//})
}

// Failed 请求失败，返回错误码信息
func Failed(code errcode.ErrCode) error {
	return status.Errorf(codes.Code(code), code.String())
}

// FailedWithMessage 请求失败，返回自定义信息
func FailedWithMessage(code errcode.ErrCode, msg string) error {
	return status.New(codes.Code(code), msg).Err()
}
