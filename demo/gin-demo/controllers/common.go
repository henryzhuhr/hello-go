package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JsonStruct 统一返回结构体
type JsonStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}

func ReturnSuccess(context *gin.Context, code int, message interface{}, data interface{}, count int64) {
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
