package orm

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type register struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	GroupId      int    `json:"group_id"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
	UserRealName string `json:"user_real_name"`
	Product      string `json:"product"`
	Email        string `json:"email"`
	ValidateType int    `json:"validate_type"`
	UserAddress  string `json:"user_address"`
	Phone        string `json:"phone"`
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error(err)
	}
	var uniqueUserName register
	db.Table("register").Where("user_name= 'yt1'").First(&uniqueUserName)
	log.Info(len(uniqueUserName.UserName))
}
