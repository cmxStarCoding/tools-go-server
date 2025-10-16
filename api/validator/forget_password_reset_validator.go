// alitools/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
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
	var request ForgetPasswordResetRequest
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
