package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-base-blog/function/model"
)

type UserApi struct{}

func (api *UserApi) Login(c *gin.Context) {
	var username = c.Query("username")
	var password = c.Query("password")
	userService.Login(username, password, c)
}

func (api *UserApi) Register(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	fmt.Println("Register:", user)
	userService.Register(user.UserName, user.Password, user.Email, c)
}

func (api *UserApi) GetUser(c *gin.Context) {
	uid := c.Query("uid")
	userService.GetUser(uid, c)
}
