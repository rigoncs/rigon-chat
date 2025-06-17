package v1

import "github.com/gin-gonic/gin"

// GetUserList 获取联系人列表
func GetUserList(c *gin.Context) {}

// LoadMyJoinedGroup 加载我加入的群组
func LoadMyJoinedGroup(c *gin.Context) {}

// GetContactInfo 获取联系人信息
func GetContactInfo(c *gin.Context) {}

// DeleteContact 删除联系人
func DeleteContact(c *gin.Context) {}

// ApplyContact 申请添加联系人
func ApplyContact(c *gin.Context) {}

// GetNewContactList 获取新的联系人申请列表
func GetNewContactList(c *gin.Context) {}

// PassContactApply 通过联系人申请
func PassContactApply(c *gin.Context) {}

// RefuseContactApply 拒绝联系人申请
func RefuseContactApply(c *gin.Context) {}

// BlackContact 拉黑联系人
func BlackContact(c *gin.Context) {}

// CancelBlackContact 取消拉黑联系人
func CancelBlackContact(c *gin.Context) {}

// GetAddGroupList 获取新的群聊申请列表
func GetAddGroupList(c *gin.Context) {}

// BlackApply 拉黑申请
func BlackApply(c *gin.Context) {}
