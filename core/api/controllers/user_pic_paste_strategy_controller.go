package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/common/utils"
	"tools/core/api/services"
	"tools/core/api/validator/pic"
)

type UserPicPasteStrategyController struct {
}

// GetUserPicPasteStrategyList 获取用户贴图策略
func (c UserPicPasteStrategyController) GetUserPicPasteStrategyList(ctx *gin.Context) {
	request, requestErr := pic.ValidateGetUserPicPasteStrategyListRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	userId := ctx.Value("UserId").(uint)
	list, listErr := services.UserPicPasteStrategyService{}.GetUserPicPasteStrategyList(request, userId)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}

	ctx.JSON(200, list)
}

// SaveUserPicPasteStrategy 保存/更新 用户贴图策略
func (c UserPicPasteStrategyController) SaveUserPicPasteStrategy(ctx *gin.Context) {
	request, requestErr := pic.ValidateSaveUserPicPasteStrategyRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	userId := ctx.Value("UserId").(uint)
	saveResult, saveResultErr := services.UserPicPasteStrategyService{}.SaveUserPicPasteStrategy(request, userId)
	if saveResultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": saveResultErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, saveResult)
}

// DeleteUserPicPasteStrategy 删除用户贴图策略
func (c UserPicPasteStrategyController) DeleteUserPicPasteStrategy(ctx *gin.Context) {
	userId := ctx.Value("UserId").(uint)
	id := utils.StringNumericToUnit(ctx.Param("id"))
	_, deleteResultErr := services.UserPicPasteStrategyService{}.DeleteUserPicPasteStrategy(id, userId)
	if deleteResultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": deleteResultErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
