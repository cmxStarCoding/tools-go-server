package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type SaveUserPicPasteStrategyRequest struct {
	ID               uint    `json:"id" form:"id" validate:"numeric" comment:"id"`
	Name             string  `json:"name" form:"name" validate:"required,min=2,max=18" comment:"贴图策略名称"`
	OriginalImageUrl string  `json:"original_image_url" form:"original_image_url" validate:"required,uri" comment:"底图"`
	StickImgUrl      string  `json:"stick_img_url" form:"stick_img_url" validate:"required" comment:"贴图地址"`
	X                uint    `json:"x" form:"x" validate:"required" comment:"x轴坐标"`
	Y                uint    `json:"y" form:"y" validate:"required" comment:"y轴坐标"`
	R                uint    `json:"r" form:"r" validate:"isRequiredR" comment:"半径"`
	Type             uint    `json:"type" form:"type" validate:"" comment:"贴图放大/缩小"`
	Multiple         float32 `json:"multiple" form:"multiple" validate:"" comment:"贴图放大缩小值"`
	BcShape          uint    `json:"bc_shape" form:"bc_shape" comment:"贴图背景区域的形状"`
	BcColor          string  `json:"bc_color" form:"bc_color" comment:"贴图背景区域的颜色"`
	SideLength       uint    `json:"side_length" form:"side_length" validate:"isRequiredSideLength" comment:"背景区域方形的边长"`
}

func ValidateSaveUserPicPasteStrategyRequest(c *gin.Context) (*SaveUserPicPasteStrategyRequest, error) {
	var request SaveUserPicPasteStrategyRequest
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
