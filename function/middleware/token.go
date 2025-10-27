package middleware

import (
	"go-base-blog/function/model"
	utilsLog "go-base-blog/function/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

// GenerateToken 生成 JWT token，返回字符串
func GenerateToken(userID uint) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		secret = []byte("your_secret_key")
	}

	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	utilsLog.FormatMessage("<<<<<<<<<< 新生成Token:", token)
	return token.SignedString(secret)
}

// JWTAuthMiddleware 返回一个 Gin 中间件，校验 Authorization: Bearer <token>
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1) 读取 Authorization header，期望格式 "Bearer <token>"
		// 2) 解析并验证 token（使用 HMAC），取出 userId 放到上下文
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		// 支持 "Bearer <token>" 或 直接传 token
		var tokenStr string
		if strings.HasPrefix(auth, "Bearer ") {
			tokenStr = strings.TrimPrefix(auth, "Bearer ")
		} else {
			tokenStr = auth
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "your_secret_key"
		}

		t, err := jwt.Parse(tokenStr, func(tk *jwt.Token) (interface{}, error) {
			// 确保签名方法为 HMAC
			if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
				model.FailWithMessage("unexpected signing method", c)
			}
			return []byte(secret), nil
		})
		if err != nil || t == nil || !t.Valid {
			model.FailWithMessage("解釋 token 失敗", c)
			return
		}

		// 解析 userId（若存在）并放到 context
		if claims, ok := t.Claims.(jwt.MapClaims); ok {
			if uid, exists := claims["userId"]; exists {
				c.Set("userId", uid)
			}
		}

		c.Next()
	}
}
