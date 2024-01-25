package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
	"tools/common/database"
	"tools/common/utils"
	"tools/core/api/models"
	"tools/core/api/validator/user"
)

var (
	ErrUserNotFound           = errors.New("未找到用户")
	EmailVerifyRecordNotFound = errors.New("邮箱验证记录未找到")
)

// UserService 用户服务
type UserService struct{}

// GetUserByID 根据用户ID获取用户信息
func (s UserService) GetUserByID(userID string) *models.UserModel {
	// 获取数据库连接
	db := database.DB
	// 调用模型方法从数据库中获取用户
	userInfo := &models.UserModel{}
	//db.Where("user_id = ?", userID).First(userInfo)
	db.First(&userInfo, userID)
	return userInfo
}

func (s UserService) UserLogin(phone string, password string) (*models.UserModel, error) {
	// 获取数据库连接
	db := database.DB
	userInfo := &models.UserModel{}
	result := db.Where("phone = ?", phone).Where("password = ?", utils.Md5Hash(password)).First(userInfo)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return userInfo, nil
}

func (s UserService) UserRegister(requestData *user.RegisterRequest) (*models.UserModel, error) {
	// 获取数据库连接
	db := database.DB
	userInfo := &models.UserFullModel{}
	err := db.Where("account = ?", requestData.Account).First(userInfo).Error

	if err == nil {
		return nil, fmt.Errorf("该账号已被注册")
	}

	var createResult = &models.UserModel{
		Account:  requestData.Account,
		Password: utils.Md5Hash(requestData.Password),
	}

	err = database.DB.Create(createResult).Error
	if err != nil {
		return nil, err
	}
	return createResult, nil
}

func (s UserService) EditUserProfile(requestData *user.EditProfileRequest, userId uint) (*models.UserModel, error) {

	userInfo := &models.UserModel{}

	result := database.DB.Where("id = ?", userId).First(userInfo)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	//修改昵称
	if requestData.Type == 1 {
		userInfo.Nickname = requestData.Nickname
		database.DB.Save(userInfo)
	}
	if requestData.Type == 2 {
		userInfo.AvatarUrl = requestData.AvatarUrl
		database.DB.Save(userInfo)
	}
	return userInfo, nil
}

func (s UserService) EditPassword(requestData *user.EditPasswordRequest, UserId uint) (string, error) {

	userInfo := &models.UserFullModel{}
	result := database.DB.Where("id = ?", UserId).Select("ID", "Password").First(userInfo)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "ok", ErrUserNotFound
	}
	oldPasswordEncrypt := utils.Md5Hash(requestData.OldPassword)
	log.Println("密码是", userInfo.Password)
	if oldPasswordEncrypt != userInfo.Password {
		return "ok", fmt.Errorf("老密码输入错误")
	}

	if requestData.NewPassword != requestData.ConfirmPassword {
		return "ok", fmt.Errorf("新密码和确认密码不一致")
	}

	if oldPasswordEncrypt == utils.Md5Hash(requestData.NewPassword) {
		return "ok", fmt.Errorf("新密码和老密码不能一致")
	}

	userInfo.Password = requestData.NewPassword
	database.DB.Save(userInfo)
	return "ok", nil
}

func (s UserService) ForgetPasswordReset(requestData *user.ForgetPasswordResetRequest) (string, error) {

	user := &models.UserFullModel{}
	result := database.DB.Where("account = ?", requestData.Account).First(user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "ok", ErrUserNotFound
	}
	if requestData.NewPassword != requestData.ConfirmPassword {
		return "ok", fmt.Errorf("新密码和确认密码不一致")
	}

	if user.Password == utils.Md5Hash(requestData.NewPassword) {
		return "ok", fmt.Errorf("新密码和老密码不能一致")
	}

	//查询最新一条邮箱验证码
	userResetPassword := &models.UserResetPasswordModel{}
	userResetPasswordResult := database.DB.Where("account = ?", requestData.Account).Last(userResetPassword)
	if userResetPasswordResult.Error != nil && errors.Is(userResetPasswordResult.Error, gorm.ErrRecordNotFound) {
		return "ok", EmailVerifyRecordNotFound
	}

	if userResetPassword.Code != requestData.EmailCode {
		return "", fmt.Errorf("邮箱验证码不正确")
	}
	user.Password = utils.Md5Hash(requestData.NewPassword)
	database.DB.Save(user)

	return "ok", nil
}

func (s UserService) SendEmailCode(requestData *user.SendEmailCodeRequest) (string, error) {

	todayStart := time.Now().Truncate(24 * time.Hour)
	var sendEmailCount int64
	countResult := database.DB.Model(&models.UserResetPasswordModel{}).Where("created_at > ?", todayStart).Count(&sendEmailCount)
	if countResult.Error != nil {
		return "", fmt.Errorf("查询统计失败")
	}
	if sendEmailCount >= 3 {
		return "", fmt.Errorf("今天发送邮件验证已达到上限")
	}

	rand.Seed(time.Now().UnixNano())
	// 生成6位数验证码
	code := rand.Intn(900000) + 100000
	smtpServer := "smtp.163.com"
	smtpPort := 25
	senderEmail := "15638276200@163.com"
	senderPassword := "QXABMJQVNQEEWQEO"
	// 收件人
	recipientEmail := requestData.UseEmail
	// 发件人昵称和邮箱地址
	senderName := "阿狸工具"
	senderAddress := senderEmail
	// 邮件内容
	subject := "重置密码验证邮件"
	body := fmt.Sprintf("<p style='color:red;'>阿狸工具：您正在重置密码，验证码为：%d</p>", code)
	// 邮件头部
	fromHeader := fmt.Sprintf("From: %s <%s>\r\n", senderName, senderAddress)
	contentType := "Content-Type: text/html; charset=UTF-8\r\n"
	message := fromHeader +
		"Subject: " + subject + "\r\n" +
		contentType +
		"\r\n" +
		body
	// 邮件服务器地址
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	// 连接到SMTP服务器
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)
	err := smtp.SendMail(smtpAddr, auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		return "", fmt.Errorf("发送邮件时发生错误")
	}
	userResetPassword := &models.UserResetPasswordModel{
		Account:  requestData.Account,
		UseEmail: requestData.UseEmail,
		Code:     strconv.Itoa(code),
	}
	database.DB.Create(userResetPassword)
	return "ok", nil
}
