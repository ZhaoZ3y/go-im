package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goim/services"
	response "goim/utils/responses"
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
