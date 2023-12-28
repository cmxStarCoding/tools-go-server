// core/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type UserLoginRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Phone    string  `json:"phone" form:"phone" validate:"required"`
	Password string  `json:"password" form:"phone" validate:"required,min=6"`
}

func ValidateUserLogin(c *gin.Context) (*UserLoginRequest, error) {
	var request UserLoginRequest
	utTrans := c.Value("Trans").(ut.Translator)

	Validate, _ := c.Get("Validate")
	validatorInstance, _ := Validate.(*validator.Validate)

	if err := c.ShouldBindJSON(&request); err != nil {
		return nil,err
	}

	//自定义错误语句
	//validatorInstance.RegisterTranslation("required", utTrans, func(ut ut.Translator) error {
	//	return ut.Add("required", "{0} 哈哈哈!", true) // see universal-translator for details
	//}, func(ut ut.Translator, fe validator.FieldError) string {
	//	t, _ := ut.T("required", fe.Field())
	//	return t
	//})

	//validatorInstance.RegisterTranslation(
	//	"email",
	//	utTrans,
	//	func(ut ut.Translator) error {
	//		return ut.Add("email", "{0}格式不正确", true)
	//	},
	//	func(ut ut.Translator, fe validator.FieldError) string {
	//		t, _ := ut.T("email", fe.Field())
	//		return t
	//	},
	//)

	// 进行进一步的验证
	err := validatorInstance.Struct(request) //这里的err是未翻译之前的
	if  err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(utTrans))
		}
		return nil,fmt.Errorf(sliceErrs[0])
	}

	return &request,nil
}