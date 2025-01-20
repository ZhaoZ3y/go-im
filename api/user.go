package api

import (
	"github.com/gin-gonic/gin"
	"goim/model"
	"goim/model/model_json"
	"goim/services"
	response "goim/utils/responses"
)

// UpdateUserAPI 更新用户信息
func UpdateUserAPI(ctx *gin.Context) {
	var user model_json.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response.ParamErr(ctx)
		return
	}

	username := ctx.GetString("username")

	User := model.User{
		UserName: username,
		NickName: user.NickName,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}

	err = services.UpdateUser(username, User)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.Success(ctx)
}

// GetUserAPI 获取用户信息
func GetUserAPI(ctx *gin.Context) {
	username := ctx.GetString("username")
	user, err := services.GetUserByUsername(username)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.OKWithData(ctx, user)
}

// SearchUserAPI 搜索用户
func SearchUserAPI(ctx *gin.Context) {
	username := ctx.Query("username")
	users, err := services.SearchUser(username)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.OKWithData(ctx, users)
}

// ChangePasswordAPI 修改密码
func ChangePasswordAPI(ctx *gin.Context) {
	var changePassword model_json.ChangePasswordReq
	err := ctx.ShouldBindJSON(&changePassword)
	if err != nil {
		response.ParamErr(ctx)
		return
	}

	username := ctx.GetString("username")
	err = services.ChangePassword(username, changePassword.OldPassword, changePassword.NewPassword)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.Success(ctx)
}
