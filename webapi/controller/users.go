package controller

import (
	"errors"

	"github.com/cm9510/cm-goweb-api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Users struct {
	base
}

// 获取用户详情
func (u Users) GetDetail(c *gin.Context) {

	id := c.Query("id")

	db := models.GetDb()

	member := new(models.Members)
	err := db.Model(&models.Members{}).
		Where(map[string]interface{}{"id": id}).
		Select([]string{"id", "name"}).
		First(&member).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.base.FailJson(c).Msg("用户不存在").Response()
		return
	} else if err != nil {
		u.base.FailJson(c).Msg(err.Error()).Response()
		return
	}

	u.base.SuccessJson(c).Data(member).Response()
}
