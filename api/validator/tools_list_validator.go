package validator

import (
	"github.com/gin-gonic/gin"
)

type GetToolListRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	CategoryId  uint `json:"category_id" form:"category_id" validate:"numeric" comment:"分类id"`
	IsRecommend uint `json:"is_recommend" form:"is_recommend" validate:"numeric" comment:"是否推荐"`
	Page        uint `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit       uint `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
}

func ValidateGetToolListRequest(c *gin.Context) (*GetToolListRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[GetToolListRequest](c)
}
