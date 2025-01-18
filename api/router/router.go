package router

import (
	"github.com/gin-gonic/gin"
	"goim/api"
)

func RouterInit() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	user := r.Group("/user")
	{
		user.POST("/register", api.RegisterAPI)
		user.POST("/login", api.LoginAPI)
		user.POST("/refresh", api.RefreshTokenAPI)
	}
	r.Run(":8080")
}
