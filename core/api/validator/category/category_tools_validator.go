package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type GetCategoryToolsRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Page  uint `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit uint `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
}

func ValidateGetCategoryToolsRequest(c *gin.Context) (*GetCategoryToolsRequest, error) {
	var request GetCategoryToolsRequest
	utTrans := c.Value("Trans").(ut.Translator)
	Validate, _ := c.Get("Validate")
	validatorInstance, _ := Validate.(*validator.Validate)
	if err := c.ShouldBindQuery(&request); err != nil {
		return nil, err
	}
	// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
	validatorInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
	// 进行进一步的验证
	err := validatorInstance.Struct(request) //这里的err是未翻译之前的
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(utTrans))
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
