package validator

import (
	"github.com/gin-gonic/gin"
)

type FeedbackRequest struct {
	ContractPhone string `json:"contract_phone" form:"contract_phone" validate:"required,gte=11,lt=12" comment:"联系电话"`
	Content       string `json:"content" form:"content" validate:"required,min=2,max=300" comment:"意见内容"`
}

func ValidateFeedbackRequest(c *gin.Context) (*FeedbackRequest, error) {
	return ValidateRequest[FeedbackRequest](c)
}
