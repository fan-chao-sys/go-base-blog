package server

import (
	"go-base-blog/function/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// Register 注册
func (us *UserService) Register(userName string, password string, email string, c *gin.Context) {
	user := model.NewUser(userName, password, email)

	// 用户密码 - 单项哈希加密(不可逆)
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashPassWord)
	err = us.db.Create(&user).Error
	if err != nil {
		lgService.Sync(success, err.Error(), userName)
		model.FailWithMessage("注册失败", c)
		return
	}
	model.OkWithData(user, c)
}

// Login 登录
func (us *UserService) Login(username string, password string, c *gin.Context) {
	u := us.GetUserName(username)
	if u.Password != password {
		lgService.Sync(fail, "登录失败", username)
		model.FailWithMessage("登录失败", c)
		return
	}
	model.Ok(c)
}

func (us *UserService) GetUserId(uid uint) *model.User {
	var user model.User
	err := us.db.Where("uid = ?", uid).Find(&user).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), user.UserName)
	}
	return &user
}

func (us *UserService) GetUser(uid string, c *gin.Context) {
	var user model.User
	us.db.Where("id = ?", uid).Find(&user)
	model.OkWithData(user, c)
}

func (us *UserService) GetUserName(username string) *model.User {
	var user model.User
	err := us.db.Where("UserName = ?", username).Find(&user).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), username)
	}
	return &user
}
