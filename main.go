package main

import (
	"fmt"
	"go-base-blog/function/initialize"
)

func main() {
	fmt.Println("🚀 启动用户管理系统...")

	initialize.DBInit()
	initialize.GinInit()
}
