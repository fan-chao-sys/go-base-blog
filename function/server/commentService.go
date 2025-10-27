package server

import (
	"github.com/gin-gonic/gin"
	"go-base-blog/function/model"
	"gorm.io/gorm"
	"strconv"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

func (cm *CommentService) CreateComment(comment model.Comment, c *gin.Context) {
	err := cm.db.Create(&comment).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), strconv.Itoa(int(comment.UserId)))
		model.FailWithMessage("创建评论失败", c)
	}
	model.OkWithData(&comment, c)
}

func (cm *CommentService) GetCommentList(postId string, c *gin.Context) {
	var comments []model.Comment
	err := cm.db.Where("post_id = ?", postId).Order("created_at desc").Find(&comments).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), "")
		model.FailWithMessage("获取评论列表失败", c)
		return
	}
	model.OkWithData(comments, c)
}
