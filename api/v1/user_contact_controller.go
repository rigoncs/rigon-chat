package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rigon-chat-server/internal/dto/request"
	"rigon-chat-server/internal/service/gorm"
	"rigon-chat-server/pkg/constants"
	"rigon-chat-server/pkg/zlog"
)

// GetUserList 获取联系人列表
func GetUserList(c *gin.Context) {
	var myUserListReq request.OwnlistRequest
	if err := c.BindJSON(&myUserListReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
	}
	message, userList, ret := gorm.UserContactService.GetUserList(myUserListReq.OwnerId)
	JSONBack(c, message, ret, userList)
}

// LoadMyJoinedGroup 加载我加入的群组
func LoadMyJoinedGroup(c *gin.Context) {
	var loadMyJoinedGroupReq request.OwnlistRequest
	if err := c.BindJSON(&loadMyJoinedGroupReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, groupList, ret := gorm.UserContactService.LoadMyJoinedGroup(loadMyJoinedGroupReq.OwnerId)
	JSONBack(c, message, ret, groupList)
}

// GetContactInfo 获取联系人信息
func GetContactInfo(c *gin.Context) {
	var getContactInfoReq request.GetContactInfoRequest
	if err := c.BindJSON(&getContactInfoReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	log.Println(getContactInfoReq)
	message, contactInfo, ret := gorm.UserContactService.GetContactInfo(getContactInfoReq.ContactId)
	JSONBack(c, message, ret, contactInfo)
}

// DeleteContact 删除联系人
func DeleteContact(c *gin.Context) {
	var deleteContactReq request.DeleteContactRequest
	if err := c.BindJSON(&deleteContactReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.DeleteContact(deleteContactReq.OwnerId, deleteContactReq.ContactId)
	JSONBack(c, message, ret, nil)
}

// ApplyContact 申请添加联系人
func ApplyContact(c *gin.Context) {
	var applyContactReq request.ApplyContactRequest
	if err := c.BindJSON(&applyContactReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.ApplyContact(applyContactReq)
	JSONBack(c, message, ret, nil)
}

// GetNewContactList 获取新的联系人申请列表
func GetNewContactList(c *gin.Context) {
	var req request.OwnlistRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, data, ret := gorm.UserContactService.GetNewContactList(req.OwnerId)
	JSONBack(c, message, ret, data)
}

// PassContactApply 通过联系人申请
func PassContactApply(c *gin.Context) {
	var passContactApplyReq request.PassContactApplyRequest
	if err := c.BindJSON(&passContactApplyReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.PassContactApply(passContactApplyReq.OwnerId, passContactApplyReq.ContactId)
	JSONBack(c, message, ret, nil)
}

// RefuseContactApply 拒绝联系人申请
func RefuseContactApply(c *gin.Context) {
	var passContactApplyReq request.PassContactApplyRequest
	if err := c.BindJSON(&passContactApplyReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.RefuseContactApply(passContactApplyReq.OwnerId, passContactApplyReq.ContactId)
	JSONBack(c, message, ret, nil)
}

// BlackContact 拉黑联系人
func BlackContact(c *gin.Context) {
	var req request.BlackContactRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.BlackContact(req.OwnerId, req.ContactId)
	JSONBack(c, message, ret, nil)
}

// CancelBlackContact 取消拉黑联系人
func CancelBlackContact(c *gin.Context) {
	var req request.BlackContactRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.CancelBlackContact(req.OwnerId, req.ContactId)
	JSONBack(c, message, ret, nil)
}

// GetAddGroupList 获取新的群聊申请列表
func GetAddGroupList(c *gin.Context) {
	var req request.AddGroupListRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, data, ret := gorm.UserContactService.GetAddGroupList(req.GroupId)
	JSONBack(c, message, ret, data)
}

// BlackApply 拉黑申请
func BlackApply(c *gin.Context) {
	var req request.BlackApplyRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.BlackApply(req.OwnerId, req.ContactId)
	JSONBack(c, message, ret, nil)
}
