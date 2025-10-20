// alitools/validator/user_validator.go

package validator

import (
	"github.com/gin-gonic/gin"
)

type EditProfileRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Type      uint   `json:"type" form:"type" validate:"required,checkTypeRange" comment:"类型"`
	Nickname  string `json:"nickname" form:"nickname" validate:"checkNickname" comment:"昵称"`
	AvatarUrl string `json:"avatar_url" form:"avatar_url" validate:"checkAvatarUrl" comment:"头像"`
}

func ValidEditProfileRequest(c *gin.Context) (*EditProfileRequest, error) {
	//调用泛型函数时可以显式指定泛型类型参数
	return ValidateRequest[EditProfileRequest](c)
}
