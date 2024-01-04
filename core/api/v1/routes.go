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
	apiV1NoNeedLogin := r.Group("/api/v1").Use(middleware.TranslationsMiddleware())
	{
		//贴图回调
		apiV1NoNeedLogin.POST("/pic_paste_notify", controllers.PicPasteController{}.Notify)
		//用户登录
		apiV1NoNeedLogin.POST("/user/login", controllers.UserController{}.UserLogin)
	}
	apiV1NeedLogin := r.Group("/api/v1").Use(middleware.TranslationsMiddleware(), middleware.JWTMiddleware())
	{
		//获取工具列表
		apiV1NeedLogin.GET("/tools_list", controllers.ToolsController{}.GetToolsList)
		//获取用户详情
		apiV1NeedLogin.GET("/user/:id", controllers.UserController{}.GetUserByID)
		//分类列表
		apiV1NeedLogin.GET("/category/list", controllers.CategoryController{}.GetCategoryList)
		//贴图服务
		apiV1NeedLogin.POST("/pic/paste", controllers.PicPasteController{}.PicPaste)
		//贴图debug
		apiV1NeedLogin.POST("/pic_paste_debug", controllers.PicPasteController{}.Debug)
		//用户使用记录
		apiV1NeedLogin.GET("/user_use_log", controllers.UserUseLogController{}.GetUserUseLogList)
	}
}
