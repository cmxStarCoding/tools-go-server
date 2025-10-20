// alitools/validator/user_validator.go

package validator

import (
	"github.com/gin-gonic/gin"
)

type ForgetPasswordResetRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Account         string `json:"account" form:"account" validate:"required" comment:"重置密码的账号"`
	NewPassword     string `json:"new_password" form:"new_password" validate:"required,min=6,max=12" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,min=6,max=12" comment:"确认密码"`
	EmailCode       string `json:"email_code" form:"email_code" validate:"required" comment:"邮箱验证码"`
}

func ValidForgetPasswordResetRequest(c *gin.Context) (*ForgetPasswordResetRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[ForgetPasswordResetRequest](c)
}
