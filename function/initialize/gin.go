package initialize

import (
	"go-base-blog/function/middleware"
	utilsLog "go-base-blog/function/utils"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func GinInit() {
	router = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Base Blog!",
		})
	})

	pu := router.Group("/public")
	pr := router.Group("/private").Use(middleware.Jwt())
	{
		pu.GET("/register", userApi.Register)
		pr.GET("/login", userApi.Login)
		pu.GET("/getPost", postApi.GetPost)
		pu.GET("/listCom", commentApi.GetCommentList)
		pr.GET("/upPost", postApi.UpdatePost)  // 本人
		pr.GET("/delPost", postApi.DeletePost) // 作者本人
	}
	{
		pr.GET("/createPost", postApi.CreatePost)      // 已認證 用戶
		pr.GET("/createCom", commentApi.CreateComment) // 已認證 用戶
	}

	router.Run(":" + os.Getenv("PORT"))
	utilsLog.FormatMessage("<<<<<<<<<<<<<<<<<<<<<<<<<< 項目啟動:端口 - %s\n <<<<<<<<<<<<<<<<<<<<<<<<<<", os.Getenv("PORT"))
}

func GetGin() *gin.Engine {
	return router
}
