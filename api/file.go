package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"goim/config"
	"goim/services"
	"goim/utils/response"
	"os"
	"strings"
)

// GetFileAPI 获取文件
func GetFileAPI(ctx *gin.Context) {
	fileName := ctx.Param("file_name")
	data, _ := os.ReadFile(config.GetConfig().Path.FilePath + fileName)
	ctx.Writer.Write(data)
}

// UploadFileAPI 上传文件
func UploadFileAPI(ctx *gin.Context) {
	namePrefix := uuid.New().String()

	userName := ctx.GetString("username")
	file, _ := ctx.FormFile("file")
	fileName := file.Filename

	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]

	newFileName := namePrefix + suffix

	ctx.SaveUploadedFile(file, config.GetConfig().Path.FilePath+newFileName)
	err := services.ChangeAvatar(userName, newFileName)
	if err != nil {
		response.InternalErr(ctx)
	}

	response.OKWithData(ctx, newFileName)
}
