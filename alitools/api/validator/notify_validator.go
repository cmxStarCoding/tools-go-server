// alitools/validator/user_validator.go

package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type NotifyRequest struct {
	BatchNo  string `json:"batch_no" form:"batch_no" validate:"required" comment:"批次号"`
	Status   uint   `json:"status" form:"status" validate:"required" comment:"状态"`
	FilePath string `json:"file_path" form:"file_path" validate:"" comment:"完成任务后的压缩包地址"`
}

func ValidateNotifyRequest(c *gin.Context) (*NotifyRequest, error) {
	var request NotifyRequest
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
