package validator

// 引入你的泛型通用校验函数和 Gin
import "github.com/gin-gonic/gin"

// GetCategoryListRequest 请求结构体
type GetCategoryListRequest struct {
	Page  uint `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit uint `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
}

// ValidateGetCategoryList 调用泛型通用函数
func ValidateGetCategoryList(c *gin.Context) (*GetCategoryListRequest, error) {
	return ValidateRequest[GetCategoryListRequest](c)
}
