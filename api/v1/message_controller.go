package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rigon-chat-server/internal/dto/request"
	"rigon-chat-server/internal/service/gorm"
	"rigon-chat-server/pkg/constants"
)

// GetMessageList 获取聊天记录
func GetMessageList(c *gin.Context) {
	var req request.GetMessageListRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, rsp, ret := gorm.MessageService.GetMessageList(req.UserOneId, req.UserTwoId)
	JSONBack(c, message, ret, rsp)
}

// GetGroupMessageList 获取群聊记录
func GetGroupMessageList(c *gin.Context) {
	var req request.GetGroupMessageListRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, rsp, ret := gorm.MessageService.GetGroupMessageList(req.GroupId)
	JSONBack(c, message, ret, rsp)
}

// UploadAvatar 上传头像
func UploadAvatar(c *gin.Context) {
	message, ret := gorm.MessageService.UploadAvatar(c)
	JSONBack(c, message, ret, nil)
}

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	message, ret := gorm.MessageService.UploadFile(c)
	JSONBack(c, message, ret, nil)
}
