package initialize

import (
	api2 "go-base-blog/api"
	server2 "go-base-blog/server"
)

var (
	userApi    = api2.UserApi{}
	postApi    = api2.PostApi{}
	commentApi = api2.CommentApi{}
)

// InitServices 初始化所有 Service 层（在 DBInit 之后调用）
func InitServices() {
	db := GetDB()

	// 先初始化 LogService（其他 Service 依赖它）
	server2.InitLogService(db)

	// 再初始化其他 Service
	userService := server2.NewUserService(db)
	postService := server2.NewPostService(db)
	comService := server2.NewCommentService(db)
	api2.SetServices(userService, postService, comService)
}
