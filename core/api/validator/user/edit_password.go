// core/validator/user_validator.go

package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type EditPasswordRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	OldPassword     string `json:"old_password" form:"old_password" validate:"required,min=6,max=12" comment:"老密码"`
	NewPassword     string `json:"new_password" form:"new_password" validate:"required,min=6,max=12" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,min=6,max=12,checkConfirmPassword" comment:"确认密码"`
}

func ValidEditPasswordRequest(c *gin.Context) (*EditPasswordRequest, error) {
	var request EditPasswordRequest
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
	////验证昵称
	validatorInstance.RegisterValidation("checkConfirmPassword", func(fl validator.FieldLevel) bool {

		//定义类型切片,1修改昵称2修改头像
		newPasswordField := fl.Parent().FieldByName("NewPassword")
		newPasswordValue := newPasswordField.String()

		confirmPasswordField := fl.Parent().FieldByName("ConfirmPassword")
		confirmPassword := confirmPasswordField.String()

		if newPasswordValue != confirmPassword {
			return false
		}

		return true
	})
	// 进行进一步的验证
	err := validatorInstance.Struct(request) //这里的err是未翻译之前的
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			if e.Tag() == "checkConfirmPassword" {
				sliceErrs = append(sliceErrs, "两次输入密码不一致")
			} else {
				sliceErrs = append(sliceErrs, e.Translate(utTrans))
			}
		}
		return nil, fmt.Errorf(sliceErrs[0])
	}

	return &request, nil
}
