package routes

import (
	"github.com/cm9510/cm-goweb-api/webapi/controller"
	"github.com/cm9510/cm-goweb-api/webapi/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRouter(route *gin.Engine) {
	// 注册路由
	r := route.Group("/api").Use(middlewares.AuthLogin())
	{
		userController := new(controller.Users)
		r.GET("/user", userController.GetDetail)
	}
}
