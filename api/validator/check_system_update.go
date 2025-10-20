package validator

import (
	"github.com/gin-gonic/gin"
)

type CheckSystemUpdateRequest struct {
	ClientVersion string `json:"client_version" form:"client_version" validate:"required" comment:"客户端版本"`
}

func ValidateCheckSystemUpdateRequest(c *gin.Context) (*CheckSystemUpdateRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[CheckSystemUpdateRequest](c)
}
