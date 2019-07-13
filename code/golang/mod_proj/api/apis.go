package apis

import (
	"net/http"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"

	"github.com/gin-gonic/gin"
	"github.com/rockdragon/daily-practice/code/golang/mod_proj/database"
)

func LoadApiRouters(prefix string, router *gin.Engine) {
	authorized := router.Group(prefix, gin.BasicAuth(gin.Accounts{
		"dev": "password",
	}))

	authorized.POST("/db/migrate", migrateDB)
	authorized.POST("/users", createUser)
	authorized.GET("/users/:username", retrieveUser)
	authorized.PUT("/users/", updateUser)
}

func updateUser(c *gin.Context) {
	var json utils.UpdateUserSpec
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	success, errors := database.UpdateUserWithProfile(json)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "DONE",
	})
}

func retrieveUser(c *gin.Context) {
	username := c.Param("username")
	success, user := database.RetrieveUserByName(username)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not exists!",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func createUser(c *gin.Context) {
	var json utils.CreateUserSpec
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	success, errors := database.CreateUserWithProfile(json)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "DONE",
	})
}

func migrateDB(c *gin.Context) {
	success, errors := database.Migrates()
	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": errors,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "DONE",
	})
}
