package api

import "go-base-blog/function/server"

var (
	userService = server.UserService{}
	postService = server.PostService{}
	comService  = server.CommentService{}
)
