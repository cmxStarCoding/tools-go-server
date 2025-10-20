package validator

import (
	"github.com/gin-gonic/gin"
)

type GetUpdateLogRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Page  uint `json:"page" form:"page" validate:"numeric" comment:"分页值"`
	Limit uint `json:"limit" form:"limit" validate:"numeric" comment:"每页数据条数"`
}

func ValidateGetUpdateLogRequest(c *gin.Context) (*GetUpdateLogRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[GetUpdateLogRequest](c)
}
