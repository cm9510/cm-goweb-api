package controller

import (
	"net/http"

	"github.com/cm9510/cm-goweb-api/internal/config"
	"github.com/gin-gonic/gin"
)

type base struct {
}

func (b *base) getConfig(c *gin.Context) *config.Items {
	conf, _ := c.Get("conf")
	items, _ := conf.(*config.Items)
	return items
}

func (b *base) SuccessJson(c *gin.Context) *JsonResponse {
	conf := b.getConfig(c)
	return &JsonResponse{
		code: conf.ApiStatus.Ok,
		msg:  conf.ApiStatus.MsgSuccess,
		data: []interface{}{},
		ctx:  c,
	}
}

func (b *base) FailJson(c *gin.Context) *JsonResponse {
	conf := b.getConfig(c)
	return &JsonResponse{
		code: conf.ApiStatus.Fail,
		msg:  conf.ApiStatus.MsgFail,
		data: []interface{}{},
		ctx:  c,
	}
}

// 展示页面
func (b *base) FetchTpl(c *gin.Context, tpl string, data ...map[string]any) {
	c.HTML(http.StatusOK, tpl, data[0])
}

// 返回json
type JsonResponse struct {
	httpState int
	code      int8
	msg       string
	data      interface{}
	ctx       *gin.Context
}

func (j *JsonResponse) HttpCode(code int) *JsonResponse {
	if code != 0 {
		j.httpState = code
	}
	return j
}

func (j *JsonResponse) Msg(msg ...string) *JsonResponse {
	if len(msg) > 0 && msg[0] != "" {
		j.msg = msg[0]
	}
	return j
}

func (j *JsonResponse) Data(data interface{}) *JsonResponse {
	j.data = data
	return j
}

func (j *JsonResponse) Response() {
	j.ctx.JSON(j.httpState, map[string]any{
		"code": j.code,
		"msg":  j.msg,
		"data": j.data,
	})
}
