// core/validator/user_validator.go

package pic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type Request struct {
	BatchNo          string  `json:"batch_no" form:"batch_no" validate:"required" comment:"批次号"`
	OriginalImageUrl string  `json:"original_image_url" form:"original_image_url" validate:"required" comment:"底图"`
	CompressFileUrl  string  `json:"compress_file_url" form:"compress_file_url" validate:"required" comment:"贴图压缩包文件地址"`
	X                uint    `json:"x" form:"x" validate:"required" comment:"x轴坐标"`
	Y                uint    `json:"y" form:"y" validate:"required" comment:"y轴坐标"`
	R                uint    `json:"r" form:"r" validate:"isSquareR" comment:"半径"`
	Type             uint    `json:"type" form:"type" validate:"required" comment:"贴图放大/缩小"`
	Multiple         float32 `json:"multiple" form:"multiple" validate:"required" comment:"贴图放大缩小值"`
	IsSquare         uint    `json:"is_square" form:"is_square" validate:"required" comment:"是否方形"`
	SideLength       uint    `json:"side_length" form:"side_length" validate:"required" comment:"方形的变长"`
}

func ValidateRequest(c *gin.Context) (*Request, error) {
	var request Request
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

	// 注册自定义验证方法
	validatorInstance.RegisterValidation("isSquareR", func(fl validator.FieldLevel) bool {
		// 获取结构体中的字段
		isSquareField := fl.Parent().FieldByName("IsSquare")
		rField := fl.Parent().FieldByName("R")

		// 获取 IsSquare 和 R 字段的值
		isSquareValue := isSquareField.Uint()
		rValue := rField.Uint()

		// 如果 IsSquare 的值等于 0，则要求 R 参数必传
		return isSquareValue == 0 && rValue > 0
	})

	// 进行进一步的验证
	err := validatorInstance.Struct(request) //这里的err是未翻译之前的
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			if e.Tag() == "isSquareR" {
				sliceErrs = append(sliceErrs, "自定义错误消息")
			} else {
				sliceErrs = append(sliceErrs, e.Translate(utTrans))
			}
		}
		return nil, fmt.Errorf(sliceErrs[0])
	}
	return &request, nil
}
