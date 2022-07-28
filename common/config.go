package common

import (
	"demon/common/common"
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	IniPath          = "common/config/app.ini"
	DefaultHost      = "127.0.0.1"
	DefaultUsername  = "root"
	DefaultPassword  = "root"
	DefaultDB        = "srun4k"
	DefaultMysqlPort = 3306
	DefaultRedisPort = 6379
	DefaultRedisPwd  = ""
	DefaultMongoPort = 27017
	DefaultMongoPwd  = ""
	DefaultLogPath   = "runtime/"
	DefaultLogLevel  = "info"
)

var (
	Config common.MainConfig
	Cfg    *ini.File
)

func init() {
	_, err := os.Stat(IniPath)
	if err != nil {
		log.Errorf("配置文件 [%s] 不存在", IniPath)
		loadDefault()
		return
	}
	Cfg, err = ini.Load(IniPath)
	if err != nil {
		log.Fatal("加载配置文件失败", err)
		return
	}
	Config.LogPath = Cfg.Section("").Key("LogPath").MustString(DefaultLogPath)
	Config.LogLevel = Cfg.Section("").Key("LogLevel").MustString(DefaultLogLevel)

	mapTo("Mysql", &Config.Mysql)
	mapTo("Mongo", &Config.Mongo)
	mapTo("Redis", &Config.Redis)
	_ = Cfg.SaveTo(IniPath)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func loadDefault() {
	// 初始化一些默认值
	Config.LogPath = DefaultLogPath
	Config.LogLevel = DefaultLogLevel
	// mysql config
	Config.Mysql.Host = DefaultHost
	Config.Mysql.Port = DefaultMysqlPort
	Config.Mysql.Username = DefaultUsername
	Config.Mysql.Password = DefaultPassword
	Config.Mysql.Dbname = DefaultDB
	// redis config
	Config.Redis.Host = DefaultHost
	Config.Redis.Port = DefaultRedisPort
	Config.Redis.Password = DefaultRedisPwd
	// mongo config
	Config.Mongo.Host = DefaultHost
	Config.Mongo.Port = DefaultMongoPort
	Config.Mongo.Password = DefaultMongoPwd

	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, &Config)
	err = cfg.SaveTo(IniPath)
	if err != nil {
		log.Errorf("映射到配置文件出错 [%v]", err)
	}
	log.Infof("配置文件 [%s] 创建成功", IniPath)
}
