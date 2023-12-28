package services

import (
	"errors"
	"gorm.io/gorm"
	"tools/common/database"
	"tools/core/api/models"
)


var (
	ErrUserNotFound = errors.New("user not found")
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

func (s UserService) UserLogin(phone string) (*models.UserModel, error) {
	// 获取数据库连接
	db := database.DB
	user := &models.UserModel{}
	result := db.Where("phone = ?",phone).First(user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound){
		return nil,ErrUserNotFound
	}
	return user,nil
}