// alitools/validator/user_validator.go

package validator

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Account  string `json:"account" form:"account" validate:"required" comment:"账号"`
	Password string `json:"password" form:"password" validate:"required,min=6" comment:"密码"`
}

func ValidateUserLogin(c *gin.Context) (*LoginRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[LoginRequest](c)
}
