package rpc

import (
	"byteDemo/cmd/user/conf"
	"byteDemo/kitex_gen/user"
	"byteDemo/kitex_gen/user/userservice"
	"context"
	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient("userservice", client.WithHostPorts(conf.Cfg.UserConf.Addr)) //由于没有使用服务注册与发现，这里需要手动绑定user服务所监听的地址
	if err != nil {
		panic(any(err))
	}
	userClient = c
}

// 实现远程调用user服务暴露的CreatUser接口
func CreatUser(ctx context.Context, req *user.Req) (*user.Resp, error) {
	resp, err := userClient.CreatUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
