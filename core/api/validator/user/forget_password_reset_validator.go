// core/validator/user_validator.go

package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
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
	utTrans := c.Value("Trans").(ut.Translator)

	Validate, _ := c.Get("Validate")
	validatorInstance, _ := Validate.(*validator.Validate)

	if err := c.ShouldBindJSON(&request); err != nil {
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

	return &request, nil
}
