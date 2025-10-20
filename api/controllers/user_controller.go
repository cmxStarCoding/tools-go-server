// service1/api/v1/controllers/user_controller.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/api/services"
	"journey/api/validator"
	"journey/models"
	"net/http"
	//"journey/api/utils"
)

// UserController 用户控制器
type UserController struct{}

// UserLogin 用户登录
func (c UserController) UserLogin(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateUserLogin,
		func(req *validator.LoginRequest) (map[string]interface{}, error) {
			return services.UserService{}.UserLogin(req.Account, req.Password)
		},
	)
}

func (c UserController) UserRegister(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateRegisterRequest,
		services.UserService{}.UserRegister,
	)
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
	userID := ctx.Value("UserId").(uint)
	// 调用用户服务获取用户信息
	userInfo := services.UserService{}.GetUserByID(userID)
	// 返回JSON数据
	ctx.JSON(200, userInfo)
}

func (c UserController) EditUserProfile(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidEditProfileRequest,
		func(req *validator.EditProfileRequest) (*models.UserModel, error) {
			userId := ctx.Value("UserId").(uint)
			return services.UserService{}.EditUserProfile(req, userId)
		},
	)
}

func (c UserController) EditUserPassword(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidEditPasswordRequest,
		func(req *validator.EditPasswordRequest) (string, error) {
			userId := ctx.Value("UserId").(uint)
			return services.UserService{}.EditPassword(req, userId)
		},
	)
}

func (c UserController) SendEmailCode(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidSendEmailCodeRequest,
		services.UserService{}.SendEmailCode,
	)
}

func (c UserController) ForgetPasswordReset(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidForgetPasswordResetRequest,
		services.UserService{}.ForgetPasswordReset,
	)
}
