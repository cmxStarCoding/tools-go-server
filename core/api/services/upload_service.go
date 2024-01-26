package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"tools/common/config"
)

type UploadService struct {
}

func (s UploadService) UploadFile(ctx *gin.Context) (string, error) {
	file, _ := ctx.FormFile("file")

	log.Println(path.Ext(file.Filename)) //文件类型
	projectConfig := config.Config
	gin.DefaultWriter.Write([]byte(""))
	dst := "../static/" + file.Filename
	// 上传文件至指定的完整文件路径
	uploadErr := ctx.SaveUploadedFile(file, dst)
	if uploadErr != nil {
		return "", fmt.Errorf(uploadErr.Error())
	}
	return projectConfig["app_domain"] + "/static/" + file.Filename, nil
}
