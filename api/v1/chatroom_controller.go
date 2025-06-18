package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rigon-chat-server/internal/dto/request"
	"rigon-chat-server/internal/service/gorm"
	"rigon-chat-server/pkg/constants"
	"rigon-chat-server/pkg/zlog"
)

// GetCurContactListInChatRoom 获取当前聊天室联系人列表
func GetCurContactListInChatRoom(c *gin.Context) {
	var req request.GetCurContactListInChatRoomRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, rspList, ret := gorm.ChatRoomService.GetCurContactListInChatRoom(req.OwnerId, req.ContactId)
	JSONBack(c, message, ret, rspList)
}
