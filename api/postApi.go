package api

import (
	"github.com/gin-gonic/gin"
	model2 "go-base-blog/model"
)

type PostApi struct{}

func (p *PostApi) CreatePost(c *gin.Context) {
	var post model2.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		model2.FailWithMessage("JSON绑定失败: "+err.Error(), c)
		return
	}
	postService.CreatePost(post, c)
}

func (p *PostApi) GetPost(c *gin.Context) {
	pid := c.Query("pid")
	postService.GetPost(pid, c)
}

func (p *PostApi) GetPostList(c *gin.Context) {
	postService.GetPostList(c)
}

func (p *PostApi) UpdatePost(c *gin.Context) {
	var post model2.Post
	c.ShouldBindJSON(&post)
	postService.UpdatePost(post, c)
}

func (p *PostApi) DeletePost(c *gin.Context) {
	pid := c.Query("pid")
	postService.DeletePost(pid, c)
}
