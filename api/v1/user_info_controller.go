package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rigon-chat-server/internal/dto/request"
	"rigon-chat-server/internal/service/gorm"
	"rigon-chat-server/pkg/constants"
	"rigon-chat-server/pkg/zlog"
)

// Register 注册
func Register(c *gin.Context) {}

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
