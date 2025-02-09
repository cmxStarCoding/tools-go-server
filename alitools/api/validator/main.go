package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"journey/common/utils"
	"reflect"
)

var Trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")
	// 注册翻译
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = zhtranslations.RegisterDefaultTranslations(v, Trans)
		// 自定义字段名翻译函数
		v.RegisterTranslation("required", Trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0}为必填字段", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", utils.ToSnakeCase(fe.Field()))
			// 将字段名转换为中文
			return t
		})

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})

		//圆形贴图的半径值验证
		v.RegisterValidation("isRequiredR", func(fl validator.FieldLevel) bool {
			// 获取结构体中的字段
			bcShapeField := fl.Parent().FieldByName("BcShape")
			rField := fl.Parent().FieldByName("R")

			// 获取 IsSquare 和 R 字段的值
			bcShapeFieldValue := bcShapeField.Uint()
			rValue := rField.Uint()
			if bcShapeFieldValue == 1 && rValue == 0 {
				return false
			}
			return true
		})

		// 方形贴图的边长值验证
		v.RegisterValidation("isRequiredSideLength", func(fl validator.FieldLevel) bool {
			// 获取结构体中的字段
			bcShapeField := fl.Parent().FieldByName("BcShape")
			sideLength := fl.Parent().FieldByName("SideLength")

			// 获取 IsSquare 和 R 字段的值
			bcShapeFieldValue := bcShapeField.Uint()
			sideLengthValue := sideLength.Uint()
			if bcShapeFieldValue == 2 && sideLengthValue == 0 {
				return false
			}
			return true
		})
		v.RegisterValidation("checkConfirmPassword", func(fl validator.FieldLevel) bool {

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

		////验证昵称
		v.RegisterValidation("checkNickname", func(fl validator.FieldLevel) bool {

			//定义类型切片,1修改昵称2修改头像
			typeField := fl.Parent().FieldByName("Type")
			typeValue := typeField.Uint()

			if typeValue == 2 {
				return true
			}

			nicknameField := fl.Parent().FieldByName("Nickname")
			nickname := nicknameField.String()
			if typeValue == 1 && nickname == "" {
				return false
			}
			return true
		})
		//验证头像
		v.RegisterValidation("checkAvatarUrl", func(fl validator.FieldLevel) bool {

			//定义类型切片,1修改昵称2修改头像
			typeField := fl.Parent().FieldByName("Type")
			typeValue := typeField.Uint()

			avatarUrlField := fl.Parent().FieldByName("AvatarUrl")
			avatarUrl := avatarUrlField.String()
			if typeValue == 2 && avatarUrl == "" {
				return false
			}
			return true
		})

		//验证type值范围
		v.RegisterValidation("checkTypeRange", func(fl validator.FieldLevel) bool {

			//定义类型切片,1修改昵称2修改头像
			typeSlice := []uint64{1, 2}
			typeField := fl.Parent().FieldByName("Type")
			typeValue := typeField.Uint()
			result := utils.ContainValue(typeSlice, typeValue)
			if result {
				return true
			}
			return false
		})

	}
}
