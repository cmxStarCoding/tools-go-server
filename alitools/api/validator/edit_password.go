// alitools/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type EditPasswordRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	OldPassword     string `json:"old_password" form:"old_password" validate:"required,min=6,max=12" comment:"老密码"`
	NewPassword     string `json:"new_password" form:"new_password" validate:"required,min=6,max=12" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,min=6,max=12,checkConfirmPassword" comment:"确认密码"`
}

func ValidEditPasswordRequest(c *gin.Context) (*EditPasswordRequest, error) {
	var request EditPasswordRequest
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
