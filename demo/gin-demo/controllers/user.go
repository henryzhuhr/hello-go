package controllers

import (
	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
// 这样做，将请求统一在一个控制器中处理，避免方法重复
type UserController struct{}

func (u UserController) GetUserInfo(context *gin.Context) {
	ReturnSuccess(context, 200, "success", "user info", 1)
}
