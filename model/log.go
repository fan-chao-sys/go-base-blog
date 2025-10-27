package model

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	ErrCode     int    `gorm:"column:err_code;type:int;not null"`
	ErrMsg      string `gorm:"column:err_msg;type:text;not null"`
	OperateName string `gorm:"column:operate_name;type:varchar(20);"`
}

func (Log) TableName() string {
	return "logs"
}

func NewLog(errCode int, msg string, operate string) *Log {
	return &Log{
		ErrCode:     errCode,
		ErrMsg:      msg,
		OperateName: operate,
	}
}
