package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtSecret = []byte("your-secret-key") // 替换为你的实际密钥

// CustomClaims 自定义的JWT声明
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID uint, nickname string) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour) // 设置Token过期时间为7天

	claims := CustomClaims{
		UserID:   userID,
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "gin-jwt-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
