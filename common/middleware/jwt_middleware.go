// core/api/controllers/middleware.go

package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"tools/core/api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware 中间件用于检测JWT Token的合法性
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is missing"})
			c.Abort()
			return
		}

		// 从Authorization Header中提取Token部分
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		// 解析Token
		claims := &utils.CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token signature"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			}
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// 将解析后的用户信息存储到Context中，以便后续的处理函数使用
		fmt.Println("请求token解析结果", claims.UserID, claims.Nickname)

		c.Set("UserId", claims.UserID)
		c.Set("Nickname", claims.Nickname)

		c.Next()
	}
}
