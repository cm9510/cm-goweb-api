package controller

import (
	"net/http"

	"github.com/cm9510/cm-goweb-api/internal/config"
	"github.com/gin-gonic/gin"
)

type base struct {
}

func (b *base) SuccessJson(c *gin.Context, args ...interface{}) {
	config := config.GetConfig()
	if msg, ok := args[0]; !ok {
		msg = "ok"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": config.ApiStatus.Ok,
		"msg":  msg,
	})
}
