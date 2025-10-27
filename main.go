package main

import (
	"fmt"
	initialize2 "go-base-blog/initialize"
)

func main() {
	fmt.Println("🚀 启动用户管理系统...")
	// 初始化 DB 与 Gin
	initialize2.DBInit()

	// 初始化 Service 层（注入数据库连接）
	initialize2.InitServices()

	// 启动 Gin（放在最后，保持原有行为）
	initialize2.GinInit()
}
