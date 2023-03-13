package controller

import (
	"github.com/gin-gonic/gin"
)

type Home struct {
	base
}

func (h *Home) Index(c *gin.Context) {

	h.base.FetchTpl(c, "home/index", map[string]any{
		"title":   "title",
		"content": "hello world",
	})
}
