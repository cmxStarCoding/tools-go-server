// service1/api/v1/routes.go

package v1

import (
	"github.com/gin-gonic/gin"
	"tools/common/middleware"
	"tools/core/api/controllers"
	//"tools/common/middleware"
)

// SetupRoutes 设置API路由
func SetupRoutes(r *gin.Engine) {

	apiv1 := r.Group("/api/v1").Use(middleware.TranslationsMiddleware())
	{
		apiv1.POST("/pic_paste_notify", controllers.PicPasteController{}.Notify)
		apiv1.POST("/user/login", controllers.UserController{}.UserLogin)
		apiv1.Use(middleware.JWTMiddleware()).GET("/user/:id", controllers.UserController{}.GetUserByID)
		apiv1.Use(middleware.JWTMiddleware()).GET("/category/list", controllers.CategoryController{}.GetCategoryList)
		apiv1.Use(middleware.JWTMiddleware()).POST("/pic/paste", controllers.PicPasteController{}.PicPaste)
	}
}
