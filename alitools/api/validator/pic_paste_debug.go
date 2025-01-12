// alitools/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type DebugRequest struct {
	OriginalImageUrl string  `json:"original_image_url" form:"original_image_url" validate:"required,uri" comment:"底图"`
	StickImgUrl      string  `json:"stick_img_url" form:"stick_img_url" validate:"required" comment:"贴图地址"`
	X                uint    `json:"x" form:"x" validate:"required" comment:"x轴坐标"`
	Y                uint    `json:"y" form:"y" validate:"required" comment:"y轴坐标"`
	R                uint    `json:"r" form:"r" validate:"isRequiredR" comment:"半径"`
	Type             uint    `json:"type" form:"type" validate:"" comment:"贴图放大/缩小"`
	Multiple         float32 `json:"multiple" form:"multiple" validate:"" comment:"贴图放大缩小值"`
	BcShape          uint    `json:"bc_shape" form:"bc_shape" comment:"贴图背景区域的形状"`
	BcColor          string  `json:"bc_color" form:"bc_color" comment:"贴图背景区域的颜色"`
	SideLength       uint    `json:"side_length" form:"side_length" validate:"isRequiredSideLength" comment:"背景区域方形的边长"`
}

func ValidateDebugRequest(c *gin.Context) (*DebugRequest, error) {
	var request DebugRequest
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
