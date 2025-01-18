package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status int    `model_json:"status"`
	Info   string `model_json:"info"`
}

var (
	// Ok 成功
	Ok = Response{
		Status: 10000,
		Info:   "OK",
	}

	// ParamError 参数错误
	ParamError = Response{
		Status: 20001,
		Info:   "Param Error",
	}

	// VerifyFailed 鉴权认证失败
	VerifyFailed = Response{
		Status: 20002,
		Info:   "Very Failed",
	}

	// UsernameOfPasswordError 登录时账号密码错误
	UsernameOfPasswordError = Response{
		Status: 20003,
		Info:   "账号或用户名错误",
	}

	// UserExist 注册时用户已存在
	UserExist = Response{
		Status: 20004,
		Info:   "用户已存在",
	}

	// InternalError 内部错误
	InternalError = Response{
		Status: 50000,
		Info:   "Internal Error",
	}

	// TokenExpired  token过期
	TokenExpired = Response{
		Status: 50001,
		Info:   "Token Expired",
	}

	// TokenHasRefreshed 该令牌在5分钟内已被刷新，无法再次刷新
	TokenHasRefreshed = Response{
		Status: 50002,
		Info:   "该令牌在5分钟内已被刷新，无法再次刷新",
	}
)

func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Ok)
}

func OKWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]any{
		"status": 10000,
		"info":   "success",
		"data":   data,
	})
}

func ParamErr(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, ParamError)
}

func VerifyErr(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, VerifyFailed)
}

func InternalErr(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, InternalError)
}

func UsernameOfPasswordErr(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, UsernameOfPasswordError)
}

func UserHasExist(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, UserExist)
}

func TokenExpiredErr(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, TokenExpired)
}

func TokenHasRefresh(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, TokenHasRefreshed)
}
