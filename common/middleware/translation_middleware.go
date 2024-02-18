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
	"sync"
)

var (
	UniMutex      sync.Mutex
	Uni           *ut.UniversalTranslator
	Validate      *validator.Validate
	ValidateMutex sync.Mutex
)

// TranslationsMiddleware 是用于处理翻译的中间件
func TranslationsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		UniMutex.Lock()
		defer UniMutex.Unlock()

		ValidateMutex.Lock()
		defer ValidateMutex.Unlock()

		//记录日志
		//gin.DefaultWriter.Write([]byte("Handling /ping request\n"))

		//多语言翻译
		if Uni == nil {
			Uni = ut.New(en.New(), zh.New())
		}
		if Validate == nil {
			Validate = validator.New()
		}

		locale := c.DefaultQuery("locale", "zh")
		trans, _ := Uni.GetTranslator(locale)

		switch locale {
		case "zh":
			zhtranslations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			entranslations.RegisterDefaultTranslations(Validate, trans)
		case "zh_tw":
			zhtwtranslations.RegisterDefaultTranslations(Validate, trans)
		default:
			zhtranslations.RegisterDefaultTranslations(Validate, trans)
		}

		//自定义错误内容
		c.Set("Validate", Validate)
		c.Set("Trans", trans)
		c.Next()
	}
}
