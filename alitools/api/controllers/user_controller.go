// service1/api/v1/controllers/user_controller.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/alitools/api/services"
	"journey/alitools/api/validator"
	"journey/common/middleware"
	"log"
	"net/http"
	"time"
	//"journey/alitools/api/utils"
)

// UserController 用户控制器
type UserController struct{}

// UserLogin 用户登录
func (c UserController) UserLogin(ctx *gin.Context) {
	// 验证用户请求参数
	request, err := validator.ValidateUserLogin(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用用户服务获取用户信息
	UserInfo, userErr := services.UserService{}.UserLogin(request.Account, request.Password)
	if userErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": userErr.Error()})
		return
	}
	jwtToken, err := middleware.GenerateToken(UserInfo.ID, UserInfo.Nickname)

	// 返回JSON数据
	ctx.JSON(200, gin.H{
		"jtw_token": jwtToken,
		"expire":    time.Now().Add(180 * 24 * time.Hour),
		"user_info": UserInfo,
	})
}

func (c UserController) UserRegister(ctx *gin.Context) {
	// 验证用户请求参数
	request, err := validator.ValidateRegisterRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用用户服务获取用户信息
	UserInfo, resultErr := services.UserService{}.UserRegister(request)
	if resultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": resultErr.Error()})
		return
	}
	jwtToken, err := middleware.GenerateToken(UserInfo.ID, UserInfo.Nickname)
	ctx.JSON(http.StatusOK, gin.H{
		"jtw_token": jwtToken,
		"expire":    time.Now().Add(7 * 24 * time.Hour),
		"user_info": UserInfo,
	})
}

func (c UserController) UserLogout(ctx *gin.Context) {
	// 调用用户服务获取用户信息
	result, err := services.UserService{}.UserLogout(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, result)
}

// GetUserByID 根据用户ID获取用户信息
func (c UserController) GetUserByID(ctx *gin.Context) {
	// 从请求参数中获取用户ID

	//fmt.Println()
	//fmt.Println(ctx.Get("Nickname"))
	userID := ctx.Value("UserId").(uint)
	// 调用用户服务获取用户信息
	userInfo := services.UserService{}.GetUserByID(userID)

	// 返回JSON数据
	ctx.JSON(200, userInfo)
}

func (c UserController) EditUserProfile(ctx *gin.Context) {
	request, err := validator.ValidEditRequest(ctx)
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
	request, err := validator.ValidEditPasswordRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := ctx.Value("UserId").(uint)
	editResult, editResultErr := services.UserService{}.EditPassword(request, userId)
	if editResultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": editResultErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, editResult)
}

func (c UserController) SendEmailCode(ctx *gin.Context) {
	request, err := validator.ValidSendEmailCodeRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("进来了")
	editResult, editResultErr := services.UserService{}.SendEmailCode(request)
	if editResultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": editResultErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, editResult)

}

func (c UserController) ForgetPasswordReset(ctx *gin.Context) {
	request, err := validator.ValidForgetPasswordResetRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	editResult, editResultErr := services.UserService{}.ForgetPasswordReset(request)
	if editResultErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": editResultErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, editResult)

}
