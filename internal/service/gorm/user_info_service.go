package gorm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"rigon-chat-server/internal/dao"
	"rigon-chat-server/internal/dto/request"
	"rigon-chat-server/internal/dto/respond"
	"rigon-chat-server/internal/model"
	"rigon-chat-server/pkg/constants"
	"rigon-chat-server/pkg/zlog"
)

type userInfoService struct{}

var UserInfoService = new(userInfoService)

// Login 登陆
func (u *userInfoService) Login(loginReq request.LoginRequest) (string, *respond.LoginRespond, int) {
	password := loginReq.Password
	var user model.UserInfo
	res := dao.GormDB.First(&user, "telephone = ?", loginReq.Telephone)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			message := "用户不存在，请注册"
			zlog.Error(message)
			return message, nil, -2
		}
		zlog.Error(res.Error.Error())
		return constants.SYSTEM_ERROR, nil, -1
	}
	if user.Password != password {
		message := "密码错误"
		zlog.Error(message)
		return message, nil, -2
	}

	loginRsp := &respond.LoginRespond{
		Uuid:      user.Uuid,
		Telephone: user.Telephone,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Birthday:  user.Birthday,
		Email:     user.Email,
		Gender:    user.Gender,
		IsAdmin:   user.IsAdmin,
		Signature: user.Signature,
		Status:    user.Status,
	}
	year, month, day := user.CreatedAt.Date()
	loginRsp.CreatedAt = fmt.Sprintf("%d-%d-%d", year, month, day)
	return "登陆成功", loginRsp, 0
}
