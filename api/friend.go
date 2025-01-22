package api

import (
	"github.com/gin-gonic/gin"
	"goim/services"
	constant "goim/utils/const"
	"goim/utils/response"
)

// GetFriendListAPI 获取好友列表
func GetFriendListAPI(ctx *gin.Context) {
	username := ctx.GetString("username")
	friends, err := services.GetFriendList(username)
	if err != nil {
		response.InternalErr(ctx)
		return
	}
	response.OKWithData(ctx, friends)
}

// AddFriendAPI 添加好友
func AddFriendAPI(ctx *gin.Context) {
	username := ctx.GetString("username")
	friendName := ctx.PostForm("friend_name")
	err := services.AddFriend(username, friendName)
	if err != nil {
		if err.Error() == constant.FriendExisted {
			response.FriendHasExisted(ctx)
			return
		}
		response.InternalErr(ctx)
		return
	}
	response.Success(ctx)
}

// DeleteFriendAPI 删除好友
func DeleteFriendAPI(ctx *gin.Context) {
	username := ctx.GetString("username")
	friendName := ctx.PostForm("friend_name")
	err := services.DeleteFriend(username, friendName)
	if err != nil {
		response.InternalErr(ctx)
		return
	}
	response.Success(ctx)
}
