// alitools/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type RegisterRequest struct {
	//Username string `json:"username" binding:"required,min=4,max=20"`
	//Phone   string  `json:"phone" form:"phone" validate:"required,email"`
	Account         string `json:"account" form:"account" validate:"required" comment:"账号"`
	Password        string `json:"password" form:"password" validate:"required,min=6" comment:"密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password" comment:"确认密码"`
}

func ValidateRegisterRequest(c *gin.Context) (*RegisterRequest, error) {
	var request RegisterRequest
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
