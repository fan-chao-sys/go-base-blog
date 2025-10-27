package api

import (
	"github.com/gin-gonic/gin"
	model2 "go-base-blog/model"
)

type CommentApi struct{}

func (cm *CommentApi) CreateComment(c *gin.Context) {
	var comment model2.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		model2.FailWithMessage("JSON绑定失败: "+err.Error(), c)
		return
	}
	comService.CreateComment(comment, c)
}

func (cm *CommentApi) GetCommentList(c *gin.Context) {
	pid := c.Query("pid")
	comService.GetCommentList(pid, c)
}
