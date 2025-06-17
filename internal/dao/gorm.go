package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rigon-chat-server/internal/config"
	"rigon-chat-server/internal/model"
	"rigon-chat-server/pkg/zlog"
)

var GormDB *gorm.DB

func init() {
	conf := config.GetConfig()
	user := conf.MysqlConfig.User
	password := conf.MysqlConfig.Password
	host := conf.MysqlConfig.Host
	port := conf.MysqlConfig.Port
	appName := conf.AppName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, appName)
	var err error
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zlog.Fatal(err.Error())
	}
	err = GormDB.AutoMigrate(&model.UserInfo{}, &model.GroupInfo{}, &model.UserContact{}, &model.Session{}, &model.ContactApply{}, &model.Message{}) // 自动迁移，如果没有建表，会自动创建对应的表
	if err != nil {
		zlog.Fatal(err.Error())
	}
}
