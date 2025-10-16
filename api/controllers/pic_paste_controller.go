package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"journey/api/services"
	"journey/api/validator"
	"journey/models"
	"net/http"
)

type PicPasteController struct{}

// PicPaste 图片贴图能力
func (c PicPasteController) PicPaste(ctx *gin.Context) {
	request, requestErr := validator.ValidateRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	userId, _ := ctx.Value("UserId").(uint)

	//创建任务
	userPicStrategy, taskIdString, err := services.UserTaskLogService{}.CreateTask(request.StrategyId, models.PicPasteMark, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.PicPasteService{}.DoTask(userPicStrategy, taskIdString, request.CompressFileUrl)
	// 返回JSON数据
	ctx.JSON(200, userPicStrategy)
}

// Notify 图片贴图python服务回调
func (c PicPasteController) Notify(ctx *gin.Context) {
	jsonData, _ := ctx.GetRawData()
	fmt.Println(string(jsonData))
	request, requestErr := validator.ValidateNotifyRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	services.UserTaskLogService{}.EditTaskStatus(request)

	// 返回JSON数据
	ctx.JSON(200, "ok")
}

// Debug  图片贴图debug接口
func (c PicPasteController) Debug(ctx *gin.Context) {
	request, requestErr := validator.ValidateDebugRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	userId, _ := ctx.Value("UserId").(uint)
	//debug
	debugResult, err := services.PicPasteService{}.Debug(request, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, debugResult)
}
