package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"content" binding:"required,min=1" gorm:"column:content;type:text"`
	PostId  uint   `json:"postid" binding:"required" gorm:"column:post_id;not null"`
	Post    Post   `json:"post,omitempty" gorm:"foreignKey:PostId;references:ID;onDelete:CASCADE" binding:"-"`
	UserId  uint   `json:"userid" binding:"required" gorm:"column:user_id;not null"`
	User    User   `json:"user,omitempty" gorm:"foreignKey:UserId" binding:"-"`
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
