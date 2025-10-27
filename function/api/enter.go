package api

import "go-base-blog/function/server"

var (
	userService *server.UserService
	postService *server.PostService
	comService  *server.CommentService
)

// SetServices 设置所有 Service 实例
func SetServices(us *server.UserService, ps *server.PostService, cs *server.CommentService) {
	userService = us
	postService = ps
	comService = cs
}
