package dao

import (
	"byteDemo/cmd/user/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var client *gorm.DB

func Init() {
	cfg := conf.Cfg.MysqlConf
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.UserName, cfg.Password, addr, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //gorm框架默认去找和结构体名字对应的复数表名，例如我们这里结构体名为user，gorm会去数据库中找users这张表，但我们在mysql中创建的表名为user，因此，我们要添加这个选项，将gorm的映射表面改为单数
		},
	})
	if err != nil {
		panic(any(err))
	}
	client = db
}
