package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type GetTaskLogListRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Page   uint `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit  uint `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
	Status uint `json:"status" form:"status" validate:"numeric" comment:"状态"`
}

func ValidateGetTaskLogListRequest(c *gin.Context) (*GetTaskLogListRequest, error) {
	var request GetTaskLogListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
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

	if request.Page == 0 {
		request.Page = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}

	return &request, nil
}
