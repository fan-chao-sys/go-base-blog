package utilgService

import (
	"fmt"
	"log"
	"time"
)

// LogInfo 记录信息日志
func LogInfo(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("[INFO] %s - %s", timestamp, message)
}

// LogError 记录错误日志
func LogError(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("[ERROR] %s - %s", timestamp, message)
}

// FormatMessage 格式化消息
func FormatMessage(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
