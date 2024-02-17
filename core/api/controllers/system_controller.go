package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/core/api/services"
	"tools/core/api/validator/system"
)

type SystemController struct {
}

func (c SystemController) FeedBack(ctx *gin.Context) {

	request, err := system.ValidateFeedbackRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := ctx.Value("UserId").(uint)
	result, resultErr := services.SystemService{}.FeedBack(request, userId)
	if resultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": resultErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, result)
}

func (c SystemController) SystemUpdateLog(ctx *gin.Context) {

	request, err := system.ValidateGetUpdateLogRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, listErr := services.SystemService{}.SystemUpdateLog(request)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, list)
}

func (c SystemController) CheckSystemUpdate(ctx *gin.Context) {
	request, err := system.ValidateCheckSystemUpdateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, listErr := services.SystemService{}.CheckSystemUpdate(request)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, list)
}

func (c SystemController) CurrentLatestVersion(ctx *gin.Context) {
	list, listErr := services.SystemService{}.CurrentLatestVersion()
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, list)
}
