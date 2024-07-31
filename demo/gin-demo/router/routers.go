package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henryzhuhr/gin-demo/controllers"
)

// Router 路由配置，返回 *gin.Engine
func Router() *gin.Engine {
	r := gin.Default()
	r.GET(
		"/hello", // 路由
		func(ctx *gin.Context) { // 处理函数
			ctx.String(http.StatusOK, "Hello World!")
		},
	)

	//支持路由分组
	user := r.Group("/user")
	{
		user.GET("/info", controllers.UserController{}.GetUserInfo)
		user.POST("/user/list", func(context *gin.Context) {
			context.String(http.StatusOK, "User List")
		})
	}
	return r
}
