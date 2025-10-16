package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/api/services"
	"journey/api/validator"
	"net/http"
)

type UserTaskLogController struct {
}

func (c UserTaskLogController) GetUserTaskLogList(ctx *gin.Context) {

	request, requestErr := validator.ValidateGetTaskLogListRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	userId := ctx.Value("UserId").(uint)
	list, listErr := services.UserTaskLogService{}.GetUserTaskLogList(request, userId)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	ctx.JSON(200, list)
}
