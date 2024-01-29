package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/common/config"
)

type UploadService struct {
}

func (s UploadService) UploadFile(ctx *gin.Context) (string, error) {
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, int64(100<<20))
	file, errMax := ctx.FormFile("file")
	if errMax != nil {
		return "", fmt.Errorf("文件最大上传允许100M")
	}
	//log.Println(path.Ext(file.Filename)) //文件类型
	projectConfig := config.Config
	//gin.DefaultWriter.Write([]byte(""))
	dst := "../static/" + file.Filename
	// 上传文件至指定的完整文件路径
	uploadErr := ctx.SaveUploadedFile(file, dst)
	if uploadErr != nil {
		return "", fmt.Errorf(uploadErr.Error())
	}
	return projectConfig["app_domain"] + "/static/" + file.Filename, nil
}
