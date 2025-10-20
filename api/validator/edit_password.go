// alitools/validator/user_validator.go

package validator

import (
	"github.com/gin-gonic/gin"
)

type EditPasswordRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	OldPassword     string `json:"old_password" form:"old_password" validate:"required,min=6,max=12" comment:"老密码"`
	NewPassword     string `json:"new_password" form:"new_password" validate:"required,min=6,max=12" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,min=6,max=12,checkConfirmPassword" comment:"确认密码"`
}

func ValidEditPasswordRequest(c *gin.Context) (*EditPasswordRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[EditPasswordRequest](c)
}
