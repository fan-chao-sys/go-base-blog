package initialize

import (
	"github.com/gin-gonic/gin"
	"go-base-blog/function/middleware"
	utilsLog "go-base-blog/function/utils"
	"os"
)

var router *gin.Engine

func GinInit() {
	utilsLog.LogInfo("<<<<<<<<<<<<<<<<<<<<<<<<<< Gin 初始化 <<<<<<<<<<<<<<<<<<<<<<<<<<")

	router = gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Base Blog!",
		})
	})

	pu := router.Group("/public")
	pr := router.Group("/private").Use(middleware.Jwt())
	{
		pu.GET("/getUser", userApi.GetUser)
		pu.POST("/register", userApi.Register)
		pu.POST("/login", userApi.Login)
		pu.GET("/getPost", postApi.GetPost)
		pu.GET("/getPostList", postApi.GetPostList)
		pu.GET("/getComList", commentApi.GetCommentList)
		pu.PUT("/upPost", postApi.UpdatePost)     // 本人
		pu.DELETE("/delPost", postApi.DeletePost) // 作者本人
	}
	{
		pr.POST("/createPost", postApi.CreatePost)      // 已認證 用戶
		pr.POST("/createCom", commentApi.CreateComment) // 已認證 用戶
	}

	err := router.Run(":8080")
	if err != nil {
		utilsLog.LogError(utilsLog.FormatMessage("<<<<<<<<<<<<<<<<<<<<<<<<<< 項目啟動失败:端口 - %s, 错误: %v\n <<<<<<<<<<<<<<<<<<<<<<<<<<", os.Getenv("PORT"), err))
		return
	}
	utilsLog.LogInfo(utilsLog.FormatMessage("<<<<<<<<<<<<<<<<<<<<<<<<<< 項目啟動成功:端口 - 8080\n <<<<<<<<<<<<<<<<<<<<<<<<<<"))
}

func GetGin() *gin.Engine {
	return router
}
