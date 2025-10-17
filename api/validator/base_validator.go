package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var Trans ut.Translator
var ValidateEngine *validator.Validate

func init() {
	ValidateEngine = validator.New()
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")
	_ = zhtranslations.RegisterDefaultTranslations(ValidateEngine, Trans)

	// 自定义字段名
	ValidateEngine.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

	// 自定义 required 翻译
	ValidateEngine.RegisterTranslation("required", Trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}为必填字段", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}

func ValidateRequest[T any](c *gin.Context) (*T, error) {
	var request T
	if err := c.ShouldBindJSON(&request); err != nil {
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
