package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/alitools/api/services"
	"journey/alitools/api/validator"
	"net/http"
)

type VipLevelController struct {
}

func (c VipLevelController) GetVipLevelList(ctx *gin.Context) {
	request, err := validator.ValidateGetVipLevelListRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, listErr := services.VipLevelService{}.GetVipLevelList(request)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, list)
}
