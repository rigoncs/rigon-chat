package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rigon-chat-server/internal/dto/request"
	"rigon-chat-server/internal/service/gorm"
	"rigon-chat-server/pkg/constants"
	"rigon-chat-server/pkg/zlog"
)

// Register 注册
func Register(c *gin.Context) {
	var registerReq request.RegisterRequest
	if err := c.BindJSON(&registerReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	fmt.Println(registerReq)
	message, userInfo, ret := gorm.UserInfoService.Register(registerReq)
	JSONBack(c, message, ret, userInfo)
}

// Login 登录
func Login(c *gin.Context) {
	var loginReq request.LoginRequest
	if err := c.BindJSON(&loginReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, userInfo, ret := gorm.UserInfoService.Login(loginReq)
	JSONBack(c, message, ret, userInfo)
}

// SmsLogin 验证码登陆
func SmsLogin(c *gin.Context) {
	var req request.SmsLoginRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, userInfo, ret := gorm.UserInfoService.SmsLogin(req)
	JSONBack(c, message, ret, userInfo)
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	var req request.UpdateUserInfoRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserInfoService.UpdateUserInfo(req)
	JSONBack(c, message, ret, nil)
}

// GetUserInfoList 获取用户列表
func GetUserInfoList(c *gin.Context) {}

// AbleUsers 启用用户
func AbleUsers(c *gin.Context) {}

// DisableUsers 禁用用户
func DisableUsers(c *gin.Context) {}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	var req request.GetUserInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, userinfo, ret := gorm.UserInfoService.GetUserInfo(req.Uuid)
	JSONBack(c, message, ret, userinfo)
}

// DeleteUsers 删除用户
func DeleteUsers(c *gin.Context) {}

// SetAdmin 设置管理员
func SetAdmin(c *gin.Context) {}

// SendSmsCode 发送短信验证码
func SendSmsCode(c *gin.Context) {
	var req request.SendSmsCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserInfoService.SendSmsCode(req.Telephone)
	JSONBack(c, message, ret, nil)
}
