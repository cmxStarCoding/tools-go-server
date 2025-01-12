package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/alitools/api/services"
	"journey/alitools/api/validator"
	"net/http"
)

type ToolsController struct {
}

func (c ToolsController) GetToolsList(ctx *gin.Context) {
	request, requestErr := validator.ValidateGetToolListRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	list, ListErr := services.ToolsService{}.GetToolsList(request)

	if ListErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ListErr.Error()})
		return
	}
	ctx.JSON(200, list)
}
