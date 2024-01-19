package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"tools/common/database"
	"tools/common/utils"
	"tools/core/api/models"
	"tools/core/api/validator/user"
)

var (
	ErrUserNotFound = errors.New("未找到用户")
)

// UserService 用户服务
type UserService struct{}

// GetUserByID 根据用户ID获取用户信息
func (s UserService) GetUserByID(userID string) *models.UserModel {
	// 获取数据库连接
	db := database.DB
	// 调用模型方法从数据库中获取用户
	user := &models.UserModel{}
	//db.Where("user_id = ?", userID).First(user)
	db.First(&user, userID)
	return user
}

func (s UserService) UserLogin(phone string, password string) (*models.UserModel, error) {
	// 获取数据库连接
	db := database.DB
	user := &models.UserModel{}
	result := db.Where("phone = ?", phone).Where("password = ?", utils.Md5Hash(password)).First(user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s UserService) EditUserProfile(requestData *user.EditProfileRequest, userId uint) (*models.UserModel, error) {

	user := &models.UserModel{}

	result := database.DB.Where("id = ?", userId).First(user)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	//修改昵称
	if requestData.Type == 1 {
		user.Nickname = requestData.Nickname
		database.DB.Save(user)
	}
	if requestData.Type == 2 {
		user.AvatarUrl = requestData.AvatarUrl
		database.DB.Save(user)
	}
	return user, nil
}

func (s UserService) EditPassword(requestData *user.EditPasswordRequest, UserId uint) (string, error) {

	user := &models.UserFullModel{}
	result := database.DB.Where("id = ?", UserId).Select("ID", "Password").First(user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "ok", ErrUserNotFound
	}
	oldPasswordEncrypt := utils.Md5Hash(requestData.OldPassword)
	log.Println("密码是", user.Password)
	if oldPasswordEncrypt != user.Password {
		return "ok", fmt.Errorf("老密码输入错误")
	}

	if requestData.NewPassword != requestData.ConfirmPassword {
		return "ok", fmt.Errorf("新密码和确认密码不一致")
	}

	if oldPasswordEncrypt == utils.Md5Hash(requestData.NewPassword) {
		return "ok", fmt.Errorf("新密码和老密码不能一致")
	}

	user.Password = requestData.NewPassword
	database.DB.Save(user)
	return "ok", nil
}
