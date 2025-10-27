package server

import (
	"github.com/gin-gonic/gin"
	"go-base-blog/middleware"
	model2 "go-base-blog/model"
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
	user := model2.NewUser(userName, password, email)

	// 用户密码 - 单项哈希加密(不可逆)
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashPassWord)
	err = us.db.Create(&user).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), userName)
		model2.FailWithMessage("注册失败", c)
		return
	}
	model2.OkWithData(user, c)
}

// Login 登录
func (us *UserService) Login(username string, password string, c *gin.Context) {
	u := us.GetUserName(username)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		lgService.Sync(fail, "登录失败", username)
		model2.FailWithMessage("登录失败", c)
		return
	}

	// 生成token
	middleware.GenerateToken(u.ID)
	model2.Ok(c)
}

func (us *UserService) GetUserId(uid uint) *model2.User {
	var user model2.User
	err := us.db.Where("uid = ?", uid).Find(&user).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), user.UserName)
	}
	return &user
}

func (us *UserService) GetUser(uid string, c *gin.Context) {
	var user model2.User
	us.db.Where("id = ?", uid).Find(&user)
	model2.OkWithData(user, c)
}

func (us *UserService) GetUserName(username string) *model2.User {
	var user model2.User
	err := us.db.Debug().Where("user_name = ?", username).Find(&user).Error
	if err != nil {
		lgService.Sync(fail, err.Error(), username)
	}
	return &user
}
