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
	api1 := r.Group("/api/v1")
	{
		api1.Use(middleware.TranslationsMiddleware())
		userController := controllers.UserController{}
		api1.POST("/user/login", userController.UserLogin)
		// 添加其他路由...
	}
	api := r.Group("/api/v1")
	{
		api.Use(middleware.JWTMiddleware(), middleware.TranslationsMiddleware())
		userController := controllers.UserController{}
		api.GET("/user/:id", userController.GetUserByID)
		// 添加其他路由...
		categoryController := controllers.CategoryController{}
		api.GET("/category/list", categoryController.GetCategoryList)
	}
}
