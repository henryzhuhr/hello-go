package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JsonStruct 统一返回结构体
type JsonStruct struct {
	Code    int         `json:"code"`  // 状态码
	Message interface{} `json:"msg"`   // 返回信息
	Data    interface{} `json:"data"`  // 返回数据
	Count   int64       `json:"count"` // 数据总数
}

func ReturnSuccess(
	context *gin.Context, // gin.Context 上下文
	code int, // 状态码
	message interface{}, // 返回信息
	data interface{}, // 返回数据
	count int64, // 数据总数
) {
	context.JSON(http.StatusOK, JsonStruct{
		Code:    code,
		Message: message,
		Data:    data,
		Count:   count,
	})

}

func ReturnError(context *gin.Context, code int, message interface{}) {
	context.JSON(http.StatusOK, JsonStruct{
		Code:    code,
		Message: message,
		Data:    nil,
		Count:   0,
	})
}
