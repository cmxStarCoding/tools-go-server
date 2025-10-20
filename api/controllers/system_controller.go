package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/api/services"
	"journey/api/validator"
	"net/http"
)

type SystemController struct {
}

func (c *SystemController) FeedBack(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateFeedbackRequest,
		func(req *validator.FeedbackRequest) (string, error) {
			UserId := ctx.Value("UserId").(uint)
			return services.SystemService{}.FeedBack(req, UserId)
		},
	)
}

func (c SystemController) SystemUpdateLog(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateGetUpdateLogRequest,
		services.SystemService{}.SystemUpdateLog,
	)
}

func (c SystemController) CheckSystemUpdate(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateCheckSystemUpdateRequest,
		services.SystemService{}.CheckSystemUpdate,
	)
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
