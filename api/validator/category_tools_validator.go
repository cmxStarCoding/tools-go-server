package validator

import "github.com/gin-gonic/gin"

// GetCategoryToolsRequest 请求结构体
type GetCategoryToolsRequest struct {
	Page          uint   `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit         uint   `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
	ClientVersion string `json:"client_version" form:"client_version" validate:"required" comment:"客户端版本"`
}

// ValidateGetCategoryToolsRequest 调用泛型通用函数
func ValidateGetCategoryToolsRequest(c *gin.Context) (*GetCategoryToolsRequest, error) {
	return ValidateRequest[GetCategoryToolsRequest](c)
}
