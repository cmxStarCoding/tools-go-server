// core/validator/user_validator.go

package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"tools/common/utils"
)

type EditProfileRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Type      uint   `json:"type" form:"type" validate:"required,checkTypeRange" comment:"类型"`
	Nickname  string `json:"nickname" form:"nickname" validate:"checkNickname,min=6,max=10" comment:"昵称"`
	AvatarUrl string `json:"avatar_url" form:"avatar_url" validate:"checkAvatarUrl" comment:"头像"`
}

func ValidEditRequest(c *gin.Context) (*EditProfileRequest, error) {
	var request EditProfileRequest
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
	validatorInstance.RegisterValidation("checkNickname", func(fl validator.FieldLevel) bool {

		//定义类型切片,1修改昵称2修改头像
		typeField := fl.Parent().FieldByName("Type")
		typeValue := typeField.Uint()

		nicknameField := fl.Parent().FieldByName("Nickname")
		nickname := nicknameField.String()
		if typeValue == 1 && nickname == "" {
			return false
		}
		return true
	})
	//验证头像
	validatorInstance.RegisterValidation("checkAvatarUrl", func(fl validator.FieldLevel) bool {

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
	validatorInstance.RegisterValidation("checkTypeRange", func(fl validator.FieldLevel) bool {

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

	// 进行进一步的验证
	err := validatorInstance.Struct(request) //这里的err是未翻译之前的
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			if e.Tag() == "checkTypeRange" {
				sliceErrs = append(sliceErrs, "type值错误")
			} else if e.Tag() == "checkNickname" {
				sliceErrs = append(sliceErrs, "nickname值不能为空")
			} else if e.Tag() == "checkAvatarUrl" {
				sliceErrs = append(sliceErrs, "avatar_url值不能为空")
			} else {
				sliceErrs = append(sliceErrs, e.Translate(utTrans))
			}
		}
		return nil, fmt.Errorf(sliceErrs[0])
	}

	return &request, nil
}
