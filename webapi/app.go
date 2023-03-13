package webapi

import (
	"fmt"

	"github.com/cm9510/cm-goweb-api/internal/config"
	"github.com/cm9510/cm-goweb-api/webapi/routes"
	"github.com/gin-gonic/gin"
)

func Start() {
	// 读取配置
	config := config.GetConfig()

	// 运行模式
	gin.SetMode(config.App.Env)

	r := gin.Default()

	// 配置入全局
	r.Use(func(ctx *gin.Context) {
		ctx.Set("conf", config)
	})

	// 注册路由
	routes.ApiRouter(r)
	routes.WebRouter(r)
	// ...

	// 开启服务
	r.Run(fmt.Sprintf(":%d", config.App.RunAt))
}
