package routes

import (
	"github.com/gin-gonic/gin"
	"journey/api/controllers"
	"journey/common/middleware"
)

// SetupRoutes 设置API路由
func SetupRoutes(r *gin.Engine) {
	//websocket
	r.GET("/ws", controllers.WebsocketController{}.MyWs)

	apiV1NoNeedLogin := r.Group("/api/v1")
	{
		//贴图回调
		apiV1NoNeedLogin.POST("/pic_paste_notify", controllers.PicPasteController{}.Notify)
		//用户登录
		apiV1NoNeedLogin.POST("/user/login", controllers.UserController{}.UserLogin)
		//用户注册
		apiV1NoNeedLogin.POST("/user/register", controllers.UserController{}.UserRegister)

		//发送邮箱验证码
		apiV1NoNeedLogin.GET("/send_email_code", controllers.UserController{}.SendEmailCode)
		//忘记密码重置
		apiV1NoNeedLogin.POST("/forget_password_reset", controllers.UserController{}.ForgetPasswordReset)
		//检测更新
		apiV1NoNeedLogin.POST("/check_system_update", controllers.SystemController{}.CheckSystemUpdate)
		//获取当前最新版本
		apiV1NoNeedLogin.GET("/current_latest_version", controllers.SystemController{}.CurrentLatestVersion)
		//获取工具列表
		apiV1NoNeedLogin.GET("/tools_list", controllers.ToolsController{}.GetToolsList)
		//获取分类工具列表
		apiV1NoNeedLogin.GET("/cate_tools_list", controllers.CategoryController{}.GetCategoryToolsList)
		//分类列表
		apiV1NoNeedLogin.GET("/category/list", controllers.CategoryController{}.GetCategoryList)
		//vip等级列表
		apiV1NoNeedLogin.GET("/vip_level_list", controllers.VipLevelController{}.GetVipLevelList)
		//更新日志
		apiV1NoNeedLogin.GET("/system_update_log", controllers.SystemController{}.SystemUpdateLog)

	}
	apiV1NeedLogin := r.Group("/api/v1").Use(middleware.JWTMiddleware())
	{
		//上传文件
		apiV1NeedLogin.POST("/upload", controllers.UploadController{}.UploadFile)
		//获取用户详情
		apiV1NeedLogin.GET("/user", controllers.UserController{}.GetUserByID)
		//apiV1NeedLogin.GET("/user/:id", controllers.UserController{}.GetUserByID)
		//用户退出登录
		apiV1NeedLogin.POST("/user/logout", controllers.UserController{}.UserLogout)
		//修改用户资料
		apiV1NeedLogin.POST("/user/edit", controllers.UserController{}.EditUserProfile)
		//修改用户密码
		apiV1NeedLogin.POST("/user/edit/password", controllers.UserController{}.EditUserPassword)
		//贴图服务
		apiV1NeedLogin.POST("/pic/paste", controllers.PicPasteController{}.PicPaste)
		//贴图debug
		apiV1NeedLogin.POST("/pic_paste_debug", controllers.PicPasteController{}.Debug)
		//用户贴图策略列表
		apiV1NeedLogin.GET("/user_pic_paste_strategy_list", controllers.UserPicPasteStrategyController{}.GetUserPicPasteStrategyList)
		//用户贴图策略 保存/更新
		apiV1NeedLogin.POST("/user_pic_paste_strategy_save", controllers.UserPicPasteStrategyController{}.SaveUserPicPasteStrategy)
		//删除用户贴图策略
		apiV1NeedLogin.DELETE("/user_pic_paste_strategy_delete/:id", controllers.UserPicPasteStrategyController{}.DeleteUserPicPasteStrategy)
		//用户使用工具记录
		apiV1NeedLogin.GET("/user_use_log", controllers.UserUseLogController{}.GetUserUseLogList)
		//用户任务列表
		apiV1NeedLogin.GET("/user_task_log", controllers.UserTaskLogController{}.GetUserTaskLogList)
		//意见反馈
		apiV1NeedLogin.POST("/feedback", controllers.SystemController{}.FeedBack)

	}
}
