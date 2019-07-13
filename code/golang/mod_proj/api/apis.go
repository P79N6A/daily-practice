package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rockdragon/daily-practice/code/golang/mod_proj/database"
)

func LoadApiRouters(prefix string, router *gin.Engine) {
	authorized := router.Group(prefix, gin.BasicAuth(gin.Accounts{
		"dev": "password",
	}))

	authorized.POST("/db/migrate", migrateDB)
}

func migrateDB(c *gin.Context) {
	conn := database.GetConnection()
	if errors := conn.AutoMigrate(&database.User{}, &database.Product{}).GetErrors(); errors != nil && len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": fmt.Sprintf("%s", errors),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "DONE",
	})
}
