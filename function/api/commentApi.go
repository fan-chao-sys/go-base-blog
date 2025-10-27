package api

import (
	"github.com/gin-gonic/gin"
	"go-base-blog/function/model"
)

type CommentApi struct{}

func (cm *CommentApi) CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		model.FailWithMessage("JSON绑定失败: "+err.Error(), c)
		return
	}
	comService.CreateComment(comment, c)
}

func (cm *CommentApi) GetCommentList(c *gin.Context) {
	pid := c.Query("pid")
	comService.GetCommentList(pid, c)
}
