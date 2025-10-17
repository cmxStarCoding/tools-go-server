package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func ValidateRequest[T any](c *gin.Context) (*T, error) {
	var request T
	if err := c.ShouldBindQuery(&request); err != nil {
		return nil, err
	}

	// 调用 validator 校验 struct tag
	if err := ValidateEngine.Struct(request); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return nil, fmt.Errorf(errs[0].Translate(Trans))
		}
		return nil, err
	}

	setDefaultPagination(&request)
	return &request, nil
}

func setDefaultPagination[T any](req *T) {
	v := reflect.ValueOf(req).Elem()
	pageField := v.FieldByName("Page")
	if pageField.IsValid() && pageField.Uint() == 0 {
		pageField.SetUint(1)
	}
	limitField := v.FieldByName("Limit")
	if limitField.IsValid() && limitField.Uint() == 0 {
		limitField.SetUint(10)
	}
}
