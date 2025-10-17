package validator

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// GetCategoryToolsRequest 请求结构体
type GetCategoryToolsRequest struct {
	Page          uint   `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit         uint   `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
	Num           uint   `json:"num" form:"num" validate:"numeric,min=3" comment:"数量"`
	ClientVersion string `json:"client_version" form:"client_version" validate:"required,NumGt2" comment:"客户端版本"`
}

// ValidateGetCategoryToolsRequest 调用泛型通用函数
func ValidateGetCategoryToolsRequest(c *gin.Context) (*GetCategoryToolsRequest, error) {

	//圆形贴图的半径值验证
	ValidateEngine.RegisterValidation("NumGt2", func(fl validator.FieldLevel) bool {
		// 获取结构体中的字段
		bcShapeField := fl.Parent().FieldByName("Num")

		// 获取 IsSquare 和 R 字段的值
		bcShapeFieldValue := bcShapeField.Uint()
		if bcShapeFieldValue < 2 {
			return false
		}
		return true
	})

	ValidateEngine.RegisterTranslation("NumGt2", Trans, func(ut ut.Translator) error {
		return ut.Add("required", "存在{0}参数，则num参数要求大于等2", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	return ValidateRequest[GetCategoryToolsRequest](c)
}
