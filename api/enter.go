package api

import (
	server2 "go-base-blog/server"
)

var (
	userService *server2.UserService
	postService *server2.PostService
	comService  *server2.CommentService
)

// SetServices 设置所有 Service 实例
func SetServices(us *server2.UserService, ps *server2.PostService, cs *server2.CommentService) {
	userService = us
	postService = ps
	comService = cs
}
