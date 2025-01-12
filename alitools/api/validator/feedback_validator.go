package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type FeedbackRequest struct {
	ContractPhone string `json:"contract_phone" form:"contract_phone" validate:"required,gte=11,lt=12" comment:"联系电话"`
	Content       string `json:"content" form:"content" validate:"required,min=2,max=300" comment:"意见内容"`
}

func ValidateFeedbackRequest(c *gin.Context) (*FeedbackRequest, error) {
	var request FeedbackRequest
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		// 参数验证失败
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return nil, fmt.Errorf(err.Error())
		}
		var sliceErrs []string
		for _, e := range errs {
			//e.Field()
			sliceErrs = append(sliceErrs, e.Translate(Trans))
		}
		return nil, fmt.Errorf(sliceErrs[0])
	}
	return &request, nil
}
