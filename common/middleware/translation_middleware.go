// middleware/translations_middleware.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	entranslations "gopkg.in/go-playground/validator.v9/translations/en"
	zhtranslations "gopkg.in/go-playground/validator.v9/translations/zh"
	zhtwtranslations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

// TranslationsMiddleware 是用于处理翻译的中间件
func TranslationsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//记录日志
		gin.DefaultWriter.Write([]byte("Handling /ping request\n"))

		//多语言翻译
		Uni = ut.New(en.New(), zh.New())
		Validate = validator.New()
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zhtranslations.RegisterDefaultTranslations(Validate, trans)
			break
		case "en":
			entranslations.RegisterDefaultTranslations(Validate, trans)
			break
		case "zh_tw":
			zhtwtranslations.RegisterDefaultTranslations(Validate, trans)
			break
		default:
			zhtranslations.RegisterDefaultTranslations(Validate, trans)
			break
		}
		//自定义错误内容
		c.Set("Validate",Validate)
		c.Set("Trans",trans)
		c.Next()
	}
}
