package apis

import "github.com/gin-gonic/gin"

func LoadApiRouters(prefix string, router *gin.Engine) {
	authorized := router.Group(prefix, gin.BasicAuth(gin.Accounts{
		"dev": "password",
	}))

	authorized.POST("/db/init", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DONE.",
		})
	})
}
