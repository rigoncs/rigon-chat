package v1

import "github.com/gin-gonic/gin"

// Register 注册
func Register(c *gin.Context) {}

// Login 登录
func Login(c *gin.Context) {}

// SmsLogin 验证码登陆
func SmsLogin(c *gin.Context) {}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(c *gin.Context) {}

// GetUserInfoList 获取用户列表
func GetUserInfoList(c *gin.Context) {}

// AbleUsers 启用用户
func AbleUsers(c *gin.Context) {}

// DisableUsers 禁用用户
func DisableUsers(c *gin.Context) {}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {}

// DeleteUsers 删除用户
func DeleteUsers(c *gin.Context) {}

// SetAdmin 设置管理员
func SetAdmin(c *gin.Context) {}

// SendSmsCode 发送短信验证码
func SendSmsCode(c *gin.Context) {}
