package api

import (
	"github.com/gin-gonic/gin"
	"goim/model"

	"goim/model/model_json"
	"goim/services"
	response "goim/utils/responses"
)

// RegisterAPI 注册
func RegisterAPI(ctx *gin.Context) {
	var user model_json.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response.ParamErr(ctx)
		return
	}

	User := model.User{
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		PassWord: user.PassWord,
		Avatar:   user.Avatar,
	}

	err = services.Register(User)
	if err != nil {
		if err.Error() == services.UserExisted {
			response.UserHasExist(ctx)
			return
		}
		response.InternalErr(ctx)
		return
	}

	response.Success(ctx)
}

// LoginAPI 登录
func LoginAPI(ctx *gin.Context) {
	var LoginReq model_json.LoginReq
	err := ctx.ShouldBindJSON(&LoginReq)
	if err != nil {
		response.ParamErr(ctx)
		return
	}

	token, err := services.Login(LoginReq.UserName, LoginReq.PassWord)
	if err != nil {
		if err.Error() == services.UserLoginErr {
			response.UsernameOfPasswordErr(ctx)
			return
		}
		response.InternalErr(ctx)
		return
	}

	response.OKWithData(ctx, token)
}

// RefreshTokenAPI 刷新token
func RefreshTokenAPI(ctx *gin.Context) {
	var refreshToken model_json.RefreshTokenReq
	err := ctx.ShouldBindJSON(&refreshToken)
	if err != nil {
		response.ParamErr(ctx)
		return
	}

	NewToken, err := services.RefreshToken(refreshToken.RefreshToken)
	if err != nil {
		if err.Error() == services.RefreshFailed {
			response.TokenHasRefresh(ctx)
			return
		}
		response.InternalErr(ctx)
		return
	}

	response.OKWithData(ctx, NewToken)
}
