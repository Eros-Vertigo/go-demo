package orm

import (
	"demon/module/orm/models/user"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error(err)
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user.User{})
	if err != nil {
		log.Error(err)
		return
	}
	temp := user.Params{Name: "test"}
	data, err := temp.Find(db)
	fmt.Println(data)
}
