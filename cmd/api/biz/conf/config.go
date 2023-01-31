package conf

import "gopkg.in/ini.v1"

type config struct {
	MysqlConf mysqlConf `ini:"mysql"`
	UserConf  userConf  `ini:"userService"` // tag要与配置文件中[]内容对应
}

type mysqlConf struct {
	UserName string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	DBName   string `ini:"dbName"`
}

type userConf struct {
	Addr string `ini:"addr"`
}

var Cfg = new(config)

const confPath = "../../common/conf/my.ini" // 读取全局的配置文件
//user kitex时再加一个../在前面

func Init() {
	err := ini.MapTo(Cfg, confPath)
	if err != nil {
		panic(any(err))
	}
}
