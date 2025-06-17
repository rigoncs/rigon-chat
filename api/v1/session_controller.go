package v1

import "github.com/gin-gonic/gin"

// OpenSession 打开会话
func OpenSession(c *gin.Context) {}

// GetUserSessionList 获取用户会话列表
func GetUserSessionList(c *gin.Context) {}

// GetGroupSessionList 获取群会话列表
func GetGroupSessionList(c *gin.Context) {}

// DeleteSession 删除会话
func DeleteSession(c *gin.Context) {}

// CheckOpenSessionAllowed 检查用户是否允许打开会话
func CheckOpenSessionAllowed(c *gin.Context) {}
