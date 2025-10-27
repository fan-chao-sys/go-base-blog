package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"column:content;type:text"`
	PostId  uint   `gorm:"column:post_id;not null"`
	Post    Post   `gorm:"foreignKey:PostId"`
	UserId  uint   `gorm:"column:user_id;not null"`
	User    User   `gorm:"foreignKey:UserId"`
}

func (Comment) TableName() string {
	return "comments"
}

func NewComment(content string, postId uint, uid uint) *Comment {
	return &Comment{
		Content: content,
		PostId:  postId,
		UserId:  uid,
	}
}
