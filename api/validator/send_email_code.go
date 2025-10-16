// alitools/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type SendEmailCodeRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Account  string `json:"account" form:"account" validate:"required,min=6,max=12" comment:"账号"`
	UseEmail string `json:"use_email" form:"use_email" validate:"required,email" comment:"验证邮箱"`
}

func ValidSendEmailCodeRequest(c *gin.Context) (*SendEmailCodeRequest, error) {
	var request SendEmailCodeRequest
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

	return &request, nil
}
