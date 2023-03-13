package routes

import (
	"github.com/cm9510/cm-goweb-api/webapi/controller"
	"github.com/gin-gonic/gin"
)

func WebRouter(route *gin.Engine) {
	// 注册网页及静态资源
	route.LoadHTMLGlob("./webapi/templates/**/*")
	route.Static("/assets", "./statics")
	// route.Static("/assets/js", "./statics/js")

	// 注册路由
	r := route.Group("/web")
	{
		homeController := new(controller.Home)
		r.GET("/home", homeController.Index)
	}
}
