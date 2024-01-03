package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/core/api/models"
	"tools/core/api/services"
	"tools/core/api/validator/pic"
)

type PicPasteController struct{}

// PicPaste 图片贴图能力
func (c PicPasteController) PicPaste(ctx *gin.Context) {
	request, requestErr := pic.ValidateRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	userId, _ := ctx.Value("UserId").(uint)
	//创建任务
	userTaskLog, err := services.UserTaskLogService{}.CreateTask(request, models.PicPasteMark, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	go services.PicPasteService{}.DoTask(request, userTaskLog["task_id"])
	// 返回JSON数据
	ctx.JSON(200, userTaskLog)
}

// Notify 图片贴图python服务回调
func (c PicPasteController) Notify(ctx *gin.Context) {
	request, requestErr := pic.ValidateNotifyRequest(ctx)
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
	request, requestErr := pic.ValidateDebugRequest(ctx)
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
