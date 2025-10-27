package middleware

import (
	"github.com/gin-gonic/gin"
)

// Jwt 返回 JWT 认证中间件
func Jwt() gin.HandlerFunc {
	// 直接返回 JWT 认证中间件函数
	return JWTAuthMiddleware()
}
