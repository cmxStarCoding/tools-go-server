package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

}

// DeleteUserPicPasteStrategy 删除用户贴图策略
func (c UserPicPasteStrategyController) DeleteUserPicPasteStrategy(ctx *gin.Context) {

}
