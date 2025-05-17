package bootstrap

import (
	"log"
	"user-service/internal/interfaces/rpc"
)

func InitRpc() {
	log.Println("启动 gRPC 服务器")
	if err := rpc.InitRpc(); err != nil {
		log.Println("初始化 gRPC 服务器失败！")
		panic(err)
	}
}
