package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name;type:varchar(20);unique;not null"`
	Password string `gorm:"column:pass_word;type:varchar(255);not null"`
	Email    string `gorm:"column:email;type:varchar(255);unique;not null"`
}

func (User) TableName() string {
	return "users"
}

func NewUser(username string, password string, email string) *User {
	return &User{
		UserName: username,
		Password: string(password),
		Email:    email,
	}
}
