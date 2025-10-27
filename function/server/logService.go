package server

import (
	"go-base-blog/function/model"
	"gorm.io/gorm"
)

type LogService struct {
	db *gorm.DB
}

func NewLogService(db *gorm.DB) *LogService {
	return &LogService{db: db}
}

func (l *LogService) Sync(errCode int, msg string, operate string) {
	log := model.NewLog(errCode, msg, operate)
	l.db.Create(&log)
}
