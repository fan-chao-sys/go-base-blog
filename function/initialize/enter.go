package initialize

import (
	"go-base-blog/function/api"
	"go-base-blog/function/server"
)

var (
	userApi    = api.UserApi{}
	postApi    = api.PostApi{}
	commentApi = api.CommentApi{}
)

// InitServices 初始化所有 Service 层（在 DBInit 之后调用）
func InitServices() {
	db := GetDB()

	// 先初始化 LogService（其他 Service 依赖它）
	server.InitLogService(db)

	// 再初始化其他 Service
	userService := server.NewUserService(db)
	postService := server.NewPostService(db)
	comService := server.NewCommentService(db)
	api.SetServices(userService, postService, comService)
}
