// alitools/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type EditProfileRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Type      uint   `json:"type" form:"type" validate:"required,checkTypeRange" comment:"类型"`
	Nickname  string `json:"nickname" form:"nickname" validate:"checkNickname" comment:"昵称"`
	AvatarUrl string `json:"avatar_url" form:"avatar_url" validate:"checkAvatarUrl" comment:"头像"`
}

func ValidEditProfileRequest(c *gin.Context) (*EditProfileRequest, error) {
	var request EditProfileRequest
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		// 参数验证失败
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return nil, fmt.Errorf(err.Error())
		}
		var sliceErrs []string
		for _, e := range errs {
			//e.Field()
			sliceErrs = append(sliceErrs, e.Translate(Trans))
		}
		return nil, fmt.Errorf(sliceErrs[0])
	}
	return &request, nil
}
