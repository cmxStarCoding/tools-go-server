package routes

import (
	"github.com/gin-gonic/gin"
	"journey/api/controllers"
	"journey/common/middleware"
	"journey/ws"
)

// SetupRoutes 设置API路由
func SetupRoutes(r *gin.Engine) {
	apiV1NoNeedLogin := r.Group("/api/v1")
	{
		category := &controllers.CategoryController{}
		//获取分类工具列表
		apiV1NoNeedLogin.GET("/cate_tools_list", category.GetCategoryToolsList)
		//分类列表
		apiV1NoNeedLogin.GET("/category/list", category.GetCategoryList)

		user := &controllers.UserController{}
		//用户登录
		apiV1NoNeedLogin.POST("/user/login", user.UserLogin)
		//用户注册
		apiV1NoNeedLogin.POST("/user/register", user.UserRegister)
		//发送邮箱验证码
		apiV1NoNeedLogin.GET("/send_email_code", user.SendEmailCode)
		//忘记密码重置
		apiV1NoNeedLogin.POST("/forget_password_reset", user.ForgetPasswordReset)

		system := &controllers.SystemController{}
		//反馈
		apiV1NoNeedLogin.GET("/feedback", system.FeedBack)
		//检测更新
		apiV1NoNeedLogin.POST("/check_system_update", system.CheckSystemUpdate)
		//获取当前最新版本
		apiV1NoNeedLogin.GET("/current_latest_version", system.CurrentLatestVersion)
		//更新日志
		apiV1NoNeedLogin.GET("/system_update_log", system.SystemUpdateLog)
	}
	apiV1NeedLogin := r.Group("/api/v1").Use(middleware.JWTMiddleware())
	{
		//websocket
		apiV1NeedLogin.GET("/ws", ws.WsHandler)
		upload := &controllers.UploadController{}
		//上传文件
		apiV1NeedLogin.POST("/upload", upload.UploadFile)

		user := &controllers.UserController{}
		//获取用户详情
		apiV1NeedLogin.GET("/user", user.GetUserByID)
		//apiV1NeedLogin.GET("/user/:id", controllers.UserController{}.GetUserByID)
		//用户退出登录
		apiV1NeedLogin.POST("/user/logout", user.UserLogout)
		//修改用户资料
		apiV1NeedLogin.POST("/user/edit", user.EditUserProfile)
		//修改用户密码
		apiV1NeedLogin.POST("/user/edit/password", user.EditUserPassword)
	}
}
