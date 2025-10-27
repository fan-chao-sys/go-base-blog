package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"column:title;type:varchar(255);not null"`
	Content string `gorm:"column:content;type:text;not null"`
	Author  uint   `gorm:"column:author;not null"`
	UserId  uint   `gorm:"column:user_id;not null;index"`
	User    User   `gorm:"foreignkey:UserID"`
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
