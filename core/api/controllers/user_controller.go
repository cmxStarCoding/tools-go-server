// service1/api/v1/controllers/user_controller.go

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/core/api/services"
	"tools/core/api/utils"
	"tools/core/api/validator"
	//"tools/core/api/utils"
)

// UserController 用户控制器
type UserController struct{}

// GetUserByID 根据用户ID获取用户信息
func (c *UserController) GetUserByID(ctx *gin.Context) {
	// 从请求参数中获取用户ID
	userID := ctx.Param("id")

	//fmt.Println(ctx.Get("UserID"))
	//fmt.Println(ctx.Get("Nickname"))

	// 调用用户服务获取用户信息
	user := services.UserService{}.GetUserByID(userID)

	// 返回JSON数据
	ctx.JSON(200, user)
}

// UserLogin 用户登录
func (c *UserController) UserLogin(ctx *gin.Context) {
	// 验证用户请求参数
	request,err := validator.ValidateUserLogin(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 调用用户服务获取用户信息
	user, userErr := services.UserService{}.UserLogin(request.Phone)
	if userErr != nil {
		// 根据不同的错误类型返回相应的 HTTP 响应
		switch userErr {
		case services.ErrUserNotFound:
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "未找到用户"})
			return
		default:
			// 处理其他错误类型，可以返回适当的 HTTP 响应或者记录日志等
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	fmt.Println(user.ID,user.Nickname)
	jwtToken,err := utils.GenerateToken(user.ID,user.Nickname)

	// 返回JSON数据
	ctx.JSON(200, gin.H{
		"jtw_token": jwtToken,
		"expire": "User created successfully",
	})
}
