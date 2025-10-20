// alitools/validator/user_validator.go

package validator

import (
	"github.com/gin-gonic/gin"
)

type SendEmailCodeRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Account  string `json:"account" form:"account" validate:"required,min=6,max=12" comment:"账号"`
	UseEmail string `json:"use_email" form:"use_email" validate:"required,email" comment:"验证邮箱"`
}

func ValidSendEmailCodeRequest(c *gin.Context) (*SendEmailCodeRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[SendEmailCodeRequest](c)
}
