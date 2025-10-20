// alitools/validator/user_validator.go

package validator

import (
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Account         string `json:"account" form:"account" validate:"required" comment:"账号"`
	Password        string `json:"password" form:"password" validate:"required,min=6" comment:"密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password" comment:"确认密码"`
}

func ValidateRegisterRequest(c *gin.Context) (*RegisterRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[RegisterRequest](c)
}
