package api

import (
	"github.com/gin-gonic/gin"
	"go-base-blog/function/model"
)

type CommentApi struct{}

func (cm *CommentApi) CreateComment(c *gin.Context) {
	var comment model.Comment
	c.ShouldBindJSON(&comment)
	comService.CreateComment(comment, c)
}

func (cm *CommentApi) GetCommentList(c *gin.Context) {
	pid := c.Query("pid")
	comService.GetCommentList(pid, c)
}
