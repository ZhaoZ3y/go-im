package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goim/model/model_json"
	"goim/services"
	"goim/utils/response"
	"strconv"
)

// GetGroupAPI 获取用户群组
func GetGroupAPI(ctx *gin.Context) {
	userId := ctx.GetUint("user_id")
	fmt.Println(userId)

	groups, err := services.GetGroup(userId)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.OKWithData(ctx, groups)
}

// CreateGroupAPI 创建群组
func CreateGroupAPI(ctx *gin.Context) {
	username := ctx.GetString("username")
	var groupJson model_json.Group
	err := ctx.ShouldBindJSON(&groupJson)
	if err != nil {
		response.ParamErr(ctx)
		return
	}

	err = services.CreateGroup(username, groupJson)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.Success(ctx)
}

// GetGroupMembersAPI 获取群成员
func GetGroupMembersAPI(ctx *gin.Context) {
	groupIdStr := ctx.Query("group_id")
	groupId, err := strconv.ParseUint(groupIdStr, 10, 64)
	if err != nil {
		response.ParamErr(ctx)
		return
	}

	members, err := services.GetGroupMembers(uint(groupId))
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.OKWithData(ctx, members)
}

// JoinGroupAPI 加入群组
func JoinGroupAPI(ctx *gin.Context) {
	username := ctx.GetString("username")
	groupUuid := ctx.PostForm("group_uuid")

	err := services.JoinGroup(username, groupUuid)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.Success(ctx)
}
