package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title" binding:"required,min=1,max=255" gorm:"column:title;type:varchar(255);not null"`
	Content string `json:"content" binding:"required,min=1" gorm:"column:content;type:text;not null"`
	Author  uint   `json:"author" binding:"required" gorm:"column:author;not null"`
	UserId  uint   `json:"userid" binding:"required" gorm:"column:user_id;not null;index"`
	User    User   `json:"user,omitempty" gorm:"foreignKey:UserId" binding:"-"`
}

func (Post) TableName() string {
	return "posts"
}

func NewPost(title string, content string, author uint, uid uint) *Post {
	return &Post{
		Title:   title,
		Content: content,
		Author:  author,
		UserId:  uid,
	}
}

// BeforeCreate 创建文章- 校验用户是否已认证
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// 查用户认证状态-校验
	return
}
