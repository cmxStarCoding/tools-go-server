package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type GetCategoryListRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Page  uint `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit uint `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
}

func ValidateGetCategoryList(c *gin.Context) (*GetCategoryListRequest, error) {
	var requestData GetCategoryListRequest
	if err := c.ShouldBindQuery(&requestData); err != nil {
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

	if requestData.Page == 0 {
		requestData.Page = 1
	}
	if requestData.Limit == 0 {
		requestData.Limit = 10
	}

	return &requestData, nil
}
