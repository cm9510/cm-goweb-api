package middlewares

import "github.com/gin-gonic/gin"

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 执行之前
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(200, gin.H{
				"msg": "未登入~",
			})
		}

		c.Next() //执行控制器

		// 执行之后
		status := c.Writer.Status()
		if status == 200 {
			// TODO...
		}
	}
}
