package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	Success = 200 // 成功状态
	Error   = 500 // 服务器错误
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		code,
		msg,
		data,
	})
}

func Fail(c *gin.Context) {
	Result(Error, "操作失败", map[string]interface{}{}, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(Error, message, map[string]interface{}{}, c)
}

func Ok(c *gin.Context) {
	Result(Success, "操作成功", map[string]interface{}{}, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, message, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, "成功", data, c)
}
