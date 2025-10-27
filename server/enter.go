package server

import (
	"go-base-blog/model"
	"gorm.io/gorm"
)

var lgService *LogService

var success = model.Success
var fail = model.Error

// InitLogService 初始化日志服务
func InitLogService(db *gorm.DB) {
	lgService = NewLogService(db)
}
