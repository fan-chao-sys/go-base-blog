package main

import (
	"fmt"
	"go-base-blog/function/initialize"
	"go-base-blog/function/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB, username, password, email string) *model.User {
	// 若已存在则返回现有用户
	var exist model.User
	if err := db.Where("user_name = ?", username).First(&exist).Error; err == nil {
		return &exist
	}

	// 生成哈希密码并创建用户
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码哈希失败:", err)
		return nil
	}
	u := model.NewUser(username, string(hash), email)
	if err := db.Create(u).Error; err != nil {
		fmt.Println("创建用户失败:", err)
		return nil
	}
	return u
}

func main() {
	fmt.Println("🚀 启动用户管理系统...")

	// 初始化 DB 与 Gin
	initialize.DBInit()

	db := initialize.GetDB()

	// 初始化两个用户及其关联数据（文章 + 评论）
	userA := createUser(db, "alice", "alicepass", "alice@example.com")
	userB := createUser(db, "bob", "bobpass", "bob@example.com")

	if userA != nil {
		postA := model.NewPost("Alice's First Post", "Hello from Alice", uint(userA.ID), uint(userA.ID))
		if err := db.Create(postA).Error; err != nil {
			fmt.Println("创建文章失败:", err)
		} else {
			commentA := model.NewComment("Nice post, Alice!", uint(postA.ID), uint(userA.ID))
			if err := db.Create(commentA).Error; err != nil {
				fmt.Println("创建评论失败:", err)
			}
		}
	}

	if userB != nil {
		postB := model.NewPost("Bob's Thoughts", "Bob says hi.", uint(userB.ID), uint(userB.ID))
		if err := db.Create(postB).Error; err != nil {
			fmt.Println("创建文章失败:", err)
		} else {
			commentB := model.NewComment("Thanks for sharing, Bob!", uint(postB.ID), uint(userB.ID))
			if err := db.Create(commentB).Error; err != nil {
				fmt.Println("创建评论失败:", err)
			}
		}
	}

	// 启动 Gin（放在最后，保持原有行为）
	initialize.GinInit()
}
