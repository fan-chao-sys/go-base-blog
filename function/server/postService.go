package server

import (
	"github.com/gin-gonic/gin"
	"go-base-blog/function/model"
	"gorm.io/gorm"
	"strconv"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (p *PostService) GetPost(postId string, c *gin.Context) {
	var post model.Post
	err := p.db.Find(&post, postId).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), strconv.Itoa(int(post.UserId)))
		model.FailWithMessage("获取文章详情失败", c)
		return
	}
	model.OkWithData(post, c)
}

func (p *PostService) GetPostList(c *gin.Context) {
	var posts []model.Post
	err := p.db.Find(&posts).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), "")
		model.FailWithMessage("获取文章列表失败", c)
		return
	}
	model.OkWithData(posts, c)
}

func (p *PostService) CreatePost(post model.Post, c *gin.Context) {
	err := p.db.Create(&post).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), strconv.Itoa(int(post.UserId)))
		model.FailWithMessage("创建文章失败: "+err.Error(), c)
		return
	}
	lgService.Sync(success, "创建文章成功", strconv.Itoa(int(post.UserId)))
	model.Ok(c)
}

func (p *PostService) UpdatePost(post model.Post, c *gin.Context) {
	err := p.db.UpdateColumns(&post).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), strconv.Itoa(int(post.UserId)))
		model.FailWithMessage("更新文章失败", c)
		return
	}
	model.OkWithData(success, c)
}

func (p *PostService) DeletePost(pid string, c *gin.Context) {
	err := p.db.Delete(&model.Post{}, pid).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), "")
		model.FailWithMessage("删除文章失败", c)
		return
	}
	model.OkWithData(success, c)
}
