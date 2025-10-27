package server

import (
	"github.com/gin-gonic/gin"
	model2 "go-base-blog/model"
	"gorm.io/gorm"
	"strconv"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

func (cm *CommentService) CreateComment(comment model2.Comment, c *gin.Context) {
	// 开启调试模式查看 SQL
	err := cm.db.Debug().Create(&comment).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), strconv.Itoa(int(comment.UserId)))
		model2.FailWithMessage("创建评论失败: "+err.Error(), c)
		return
	}
	model2.Ok(c)
}

func (cm *CommentService) GetCommentList(postId string, c *gin.Context) {
	var comments []model2.Comment
	err := cm.db.Where("post_id = ?", postId).Order("created_at desc").Find(&comments).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), "")
		model2.FailWithMessage("获取评论列表失败", c)
		return
	}
	model2.OkWithData(comments, c)
}
