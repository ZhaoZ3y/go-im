package router

import (
	"github.com/gin-gonic/gin"
	"goim/api"
	"goim/middleware"
)

func RouterInit() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", api.RegisterAPI)
	login := r.Group("/login")
	{
		login.POST("", api.LoginAPI)
		login.POST("/refresh", api.RefreshTokenAPI)
	}

	user := r.Group("/user")
	user.Use(middleware.JwtMiddleware())
	{
		user.PUT("/update", api.UpdateUserAPI)
		user.GET("/info", api.GetUserAPI)
		user.GET("/search", api.SearchUserAPI)
		user.PUT("/changePassword", api.ChangePasswordAPI)
	}

	group := r.Group("/group")
	group.Use(middleware.JwtMiddleware())
	{
		group.GET("/list", api.GetGroupAPI)
	}

	r.Run(":8080")
}
