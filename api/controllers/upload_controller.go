package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/api/services"
	"net/http"
)

type UploadController struct {
}

func (c UploadController) UploadFile(ctx *gin.Context) {

	path, err := services.UploadService{}.UploadFile(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 单文件
	ctx.JSON(http.StatusOK, gin.H{
		"path": path,
	})
}
