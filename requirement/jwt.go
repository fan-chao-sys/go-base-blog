// package requirement
//
// import (
//
//	"fmt"
//	"github.com/golang-jwt/jwt/v5"
//	"time"
//
// )
//
// // CustomClaims 自定义声明（可存储用户ID、角色等信息）
//
//	type CustomClaims struct {
//		UserID int    `json:"user_id"`
//		Role   string `json:"role"`
package requirement

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义声明（可存储用户ID、角色等信息）
type CustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID int, role string, secret []byte) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 令牌有效期24小时
			Issuer:    "your-app",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名密钥（实际项目中需安全存储，如环境变量）
	signedToken, err := token.SignedString(secret)
	return signedToken, err
}

// ValidateToken 验证JWT令牌
func ValidateToken(tokenString string, secret []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

//	}
