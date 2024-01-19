// service1/api/v1/controllers/user_controller.go

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tools/core/api/services"
	"tools/core/api/utils"
	"tools/core/api/validator/user"
	//"tools/core/api/utils"
)

// UserController 用户控制器
type UserController struct{}

// UserLogin 用户登录
func (c UserController) UserLogin(ctx *gin.Context) {
	// 验证用户请求参数
	request, err := user.ValidateUserLogin(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用用户服务获取用户信息
	UserInfo, userErr := services.UserService{}.UserLogin(request.Phone, request.Password)
	if userErr != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": userErr.Error()})
		return

		// 根据不同的错误类型返回相应的 HTTP 响应
		//switch userErr {
		//case services.ErrUserNotFound:
		//	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "未找到用户"})
		//	return
		//default:
		//	// 处理其他错误类型，可以返回适当的 HTTP 响应或者记录日志等
		//	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		//	return
		//}
	}
	fmt.Println(UserInfo.ID, UserInfo.Nickname)
	jwtToken, err := utils.GenerateToken(UserInfo.ID, UserInfo.Nickname)

	// 返回JSON数据
	ctx.JSON(200, gin.H{
		"jtw_token": jwtToken,
		"expire":    time.Now().Add(7 * 24 * time.Hour),
		"user_info": UserInfo,
	})
}

// GetUserByID 根据用户ID获取用户信息
func (c UserController) GetUserByID(ctx *gin.Context) {
	// 从请求参数中获取用户ID
	userID := ctx.Param("id")

	//fmt.Println(ctx.Get("UserID"))
	//fmt.Println(ctx.Get("Nickname"))

	// 调用用户服务获取用户信息
	userInfo := services.UserService{}.GetUserByID(userID)

	// 返回JSON数据
	ctx.JSON(200, userInfo)
}

func (c UserController) EditUserProfile(ctx *gin.Context) {
	request, err := user.ValidEditRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := ctx.Value("UserId").(uint)
	result, resultErr := services.UserService{}.EditUserProfile(request, userId)
	if resultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": resultErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c UserController) EditUserPassword(ctx *gin.Context) {

}
