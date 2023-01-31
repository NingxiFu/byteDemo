package main

import (
	"byteDemo/cmd/user/conf"
	"byteDemo/cmd/user/dao"
	user "byteDemo/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	conf.Init()                                                    //读取配置文件
	dao.Init()                                                     // 初始化数据库
	addr, err := net.ResolveTCPAddr("tcp", conf.Cfg.UserConf.Addr) //绑定服务地址，如果使用服务注册，可将本行替换成服务注册配置
	if err != nil {
		panic(any(err))
	}

	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr)) // 建立一个服务端

	err = svr.Run() // 启动服务

	if err != nil {
		log.Println(err.Error())
	}
}
