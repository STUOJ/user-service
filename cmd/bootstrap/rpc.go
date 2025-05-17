package bootstrap

import (
	"log"
	"user-service/internal/infrastructure/nacos"
	"user-service/internal/interfaces/rpc"
)

func InitRpc() {
	if err := nacos.NacosClient.Register(); err != nil {
		panic(err)
	}

	log.Println("启动 gRPC 服务器")
	if err := rpc.Init(); err != nil {
		log.Println("初始化 gRPC 服务器失败！")
		panic(err)
	}
}
