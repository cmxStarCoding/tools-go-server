package controllers

import (
	"journey/api/services"
	"journey/api/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
}

func (c *SystemController) FeedBack(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateFeedbackRequest,
		func(ctx *gin.Context, req *validator.FeedbackRequest) (string, error) {
			UserId := ctx.Value("UserId").(uint)
			return services.SystemService{}.FeedBack(req, UserId)
		},
	)
}

func (c SystemController) SystemUpdateLog(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateGetUpdateLogRequest,
		func(ctx *gin.Context, req *validator.GetUpdateLogRequest) (map[string]interface{}, error) {
			return services.SystemService{}.SystemUpdateLog(req)
		},
	)
}

func (c SystemController) CheckSystemUpdate(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateCheckSystemUpdateRequest,
		func(ctx *gin.Context, req *validator.CheckSystemUpdateRequest) (map[string]any, error) {
			return services.SystemService{}.CheckSystemUpdate(req)
		},
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
