package common

// 公共配置结构体

type MainConfig struct {
	LogPath  string `comment:"日志路径"`
	LogLevel string `comment:"日志登记"`

	Mysql mysqlConfig `comment:"mysql数据库配置"`

	Redis redisConfig `comment:"redis数据库配置"`

	Mongo mongoConfig `comment:"mongo数据库配置"`
}

type mysqlConfig struct {
	Host     string `ini:"host" comment:"host"`
	Port     int    `ini:"port" comment:"端口"`
	Username string `ini:"username" comment:"用户名"`
	Password string `ini:"password" comment:"密码"`
	Dbname   string `ini:"dbname" comment:"数据库名称"`
}

type redisConfig struct {
	Host     string `ini:"host" comment:"host"`
	Port     int    `ini:"port" comment:"端口"`
	Password string `ini:"password" comment:"密码"`
}

type mongoConfig struct {
	Host     string `ini:"host" comment:"host"`
	Port     int    `ini:"port" comment:"端口"`
	Password string `ini:"password" comment:"密码"`
}
