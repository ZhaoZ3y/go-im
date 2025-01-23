package api

import (
	"github.com/gin-gonic/gin"
	"goim/model"
	"goim/services"
	"goim/utils/response"
)

// GetMessages 获取消息
func GetMessages(ctx *gin.Context) {
	var messageReq model.MessageReq
	err := ctx.BindJSON(&messageReq)
	if err != nil {
		response.ParamErr(ctx)
		return
	}
	username := ctx.GetString("username")

	messages, err := services.GetMessages(username, messageReq)
	if err != nil {
		response.InternalErr(ctx)
		return
	}

	response.OKWithData(ctx, messages)
}
