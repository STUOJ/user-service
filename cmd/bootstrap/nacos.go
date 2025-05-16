package bootstrap

import (
	"log"
	"user-service/internal/infrastructure/nacos"
)

func InitNacos() {
	err := nacos.InitNacos()
	if err != nil {
		log.Println("初始化 Nacos 失败！")
		panic(err)
	}
	log.Println("初始化 Nacos 成功")

	err = nacos.LoadConfig()
	if err != nil {
		log.Println("加载 Nacos 配置失败！")
		panic(err)
	}
	log.Println("加载 Nacos 配置成功")
}
