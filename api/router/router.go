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

	//用户登录注册API
	r.POST("/register", api.RegisterAPI)
	login := r.Group("/login")
	{
		login.POST("", api.LoginAPI)
		login.POST("/refresh", api.RefreshTokenAPI)
	}

	//用户信息API
	user := r.Group("/user")
	user.Use(middleware.JwtMiddleware())
	{
		user.PUT("/update", api.UpdateUserAPI)
		user.GET("/info", api.GetUserAPI)
		user.GET("/search", api.SearchUserAPI)
		user.PUT("/changePassword", api.ChangePasswordAPI)
	}

	//群组相关API
	group := r.Group("/group")
	group.Use(middleware.JwtMiddleware())
	{
		group.GET("/list", api.GetGroupAPI)
		group.POST("/create", api.CreateGroupAPI)
		group.GET("/members", api.GetGroupMembersAPI)
		group.POST("/join", api.JoinGroupAPI)
		group.DELETE("/quit", api.ExitGroupAPI)
	}

	//好友相关API
	friend := r.Group("/friend")
	friend.Use(middleware.JwtMiddleware())
	{
		friend.GET("/list", api.GetFriendListAPI)
		friend.POST("/add", api.AddFriendAPI)
		friend.DELETE("/delete", api.DeleteFriendAPI)
	}

	r.Run(":8080")
}
