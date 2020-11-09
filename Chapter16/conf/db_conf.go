package conf

type DBConf struct {
	Driver string
	// Host 主机地址
	Host string
	// Port 主机端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// DbName 数据库名称
	DbName string
	// Charset 数据库编码
	Charset string
}
