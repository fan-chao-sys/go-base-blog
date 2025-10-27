package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mw before")
		JWTAuthMiddleware()
		c.Next() // 调用下个中间件,如没有，及调用接口本身
		fmt.Println("mw after")
	}
}
