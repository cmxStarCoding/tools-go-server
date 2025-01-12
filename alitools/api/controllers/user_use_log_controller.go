package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/alitools/api/services"
	"journey/alitools/api/validator"
	"net/http"
)

// UserUseLogController 分类控制器
type UserUseLogController struct{}

// GetUserUseLogList 获取分类列表
func (c UserUseLogController) GetUserUseLogList(ctx *gin.Context) {
	request, err := validator.ValidateGetUserUseLogListRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := ctx.Value("UserId").(uint)
	list, listErr := services.UserUseLogService{}.UserUseLogList(request, userId)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, list)
}
