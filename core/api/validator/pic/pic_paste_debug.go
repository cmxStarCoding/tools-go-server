// core/validator/user_validator.go

package pic

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
	Type             uint    `json:"type" form:"type" validate:"required" comment:"贴图放大/缩小"`
	Multiple         float32 `json:"multiple" form:"multiple" validate:"required" comment:"贴图放大缩小值"`
	IsSquare         uint    `json:"is_square" form:"is_square" comment:"是否方形"`
	SideLength       uint    `json:"side_length" form:"side_length" validate:"isRequiredSideLength" comment:"方形的变长"`
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
	//圆形贴图的半径值验证
	validatorInstance.RegisterValidation("isRequiredR", func(fl validator.FieldLevel) bool {
		// 获取结构体中的字段
		isSquareField := fl.Parent().FieldByName("IsSquare")
		rField := fl.Parent().FieldByName("R")

		// 获取 IsSquare 和 R 字段的值
		isSquareValue := isSquareField.Uint()
		rValue := rField.Uint()
		if isSquareValue == 0 && rValue == 0 {
			return false
		}
		return true
	})

	// 方形贴图的边长值验证
	validatorInstance.RegisterValidation("isRequiredSideLength", func(fl validator.FieldLevel) bool {
		// 获取结构体中的字段
		isSquareField := fl.Parent().FieldByName("IsSquare")
		sideLength := fl.Parent().FieldByName("SideLength")

		// 获取 IsSquare 和 R 字段的值
		isSquareValue := isSquareField.Uint()
		sideLengthValue := sideLength.Uint()
		if isSquareValue == 1 && sideLengthValue == 0 {
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
			sliceErrs = append(sliceErrs, e.Translate(utTrans))
		}
		return nil, fmt.Errorf(sliceErrs[0])
	}
	return &request, nil
}
