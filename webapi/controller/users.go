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

// func (u *Users)NewUsers() *Users {
// 	return &Users{}
// }

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
		u.base.SuccessJson(c)
		// c.JSON(200, gin.H{
		// 	"err": "用户不存在",
		// })
		return
	} else if err != nil {
		c.JSON(200, gin.H{
			"err": "11111 " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": member,
	})
}
