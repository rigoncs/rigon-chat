package v1

import "github.com/gin-gonic/gin"

// CreateGroup 创建群组
func CreateGroup(c *gin.Context) {}

// LoadMyGroup 获取我创建的群聊
func LoadMyGroup(c *gin.Context) {}

// CheckGroupAddMode 检查群组加入模式
func CheckGroupAddMode(c *gin.Context) {}

// EnterGroupDirectly 直接加入群组
func EnterGroupDirectly(c *gin.Context) {}

// LeaveGroup 退出群组
func LeaveGroup(c *gin.Context) {}

// DismissGroup 解散群组
func DismissGroup(c *gin.Context) {}

// GetGroupInfo 获取群组信息
func GetGroupInfo(c *gin.Context) {}

// GetGroupInfoList 获取群组列表 - 管理员
func GetGroupInfoList(c *gin.Context) {}

// DeleteGroups 删除群组 - 管理员
func DeleteGroups(c *gin.Context) {}

// SetGroupsStatus 设置群组是否启用
func SetGroupsStatus(c *gin.Context) {}

// UpdateGroupInfo 更新群组信息
func UpdateGroupInfo(c *gin.Context) {}

// GetGroupMemberList 获取群组成员列表
func GetGroupMemberList(c *gin.Context) {}

// RemoveGroupMembers 移除群组成员
func RemoveGroupMembers(c *gin.Context) {}
