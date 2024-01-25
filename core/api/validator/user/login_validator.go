// core/validator/user_validator.go

package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type LoginRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Account  string `json:"account" form:"account" validate:"required" comment:"账号"`
	Password string `json:"password" form:"password" validate:"required,min=6" comment:"密码"`
}

func ValidateUserLogin(c *gin.Context) (*LoginRequest, error) {
	var request LoginRequest
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
