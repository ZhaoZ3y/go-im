package router

import (
	"github.com/gin-gonic/gin"
	"goim/api"
	"goim/middleware"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	socket := RunSocket
	r.GET("/ws", socket)

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
		user.GET("/detail", api.GetUserInfoAPI)
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

	//消息相关API
	message := r.Group("/message")
	message.Use(middleware.JwtMiddleware())
	{
		message.GET("", api.GetMessages)
	}

	//文件相关API
	file := r.Group("/file")
	file.Use(middleware.JwtMiddleware())
	{
		file.GET("/:fileName", api.GetFileAPI)
		file.POST("", api.UploadFileAPI)
	}

	return r
}
