package mysql

// mysql 数据库

import (
	"demon/common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		common.Config.Mysql.Username,
		common.Config.Mysql.Password,
		common.Config.Mysql.Host,
		common.Config.Mysql.Port,
		common.Config.Mysql.Dbname)

	Db, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		log.Fatalf("mysql.init err : [%v]", err)
	} else {
		log.Infof("Mysql [%s:%d] 初始化完成", common.Config.Mysql.Host, common.Config.Mysql.Port)
	}
}