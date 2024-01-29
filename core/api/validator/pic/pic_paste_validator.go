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
	//OriginalImageUrl string           `json:"original_image_url" form:"original_image_url" validate:"required,uri" comment:"底图"`
	CompressFileUrl string `json:"compress_file_url" form:"compress_file_url" validate:"required" comment:"贴图压缩包文件地址"`
	StrategyId      uint   `json:"strategy_id" form:"strategy_id" validate:"required" comment:"策略id"`

	//X                uint             `json:"x" form:"x" validate:"required" comment:"x轴坐标"`
	//Y                uint             `json:"y" form:"y" validate:"required" comment:"y轴坐标"`
	//R                uint             `json:"r" form:"r" validate:"isRequiredR" comment:"半径"`
	//Type             uint             `json:"type" form:"type" validate:"required" comment:"贴图放大/缩小"`
	//Multiple         float32          `json:"multiple" form:"multiple" validate:"required" comment:"贴图放大缩小值"`
	//BcShape          uint             `json:"bc_shape" form:"bc_shape" comment:"贴图背景区域的形状"`
	//BcColor          string           `json:"bc_color" form:"bc_color" comment:"贴图背景区域的颜色"`
	//SideLength       uint             `json:"side_length" form:"side_length" validate:"isRequiredSideLength" comment:"背景区域方形的边长"`
	//Aa               map[string]any   `json:"aa" form:"aa" comment:"字典测试"`
	//Da               []map[string]any `json:"dd" form:"dd" comment:"切片测试"`
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
