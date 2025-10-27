package api

import (
	"github.com/gin-gonic/gin"
	"go-base-blog/function/model"
)

type UserApi struct{}

func (api *UserApi) Login(c *gin.Context) {
	var username = c.Param("username")
	var password = c.Param("password")
	userService.Login(username, password, c)
}

func (api *UserApi) Register(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	userService.Register(user.UserName, user.Password, user.Email, c)
}
