package bootstrap

import (
	"log"
	"user-service/pkg/config"
)

func InitConfig() {
	err := config.InitConfig()
	if err != nil {
		log.Println("初始化配置失败！")
		panic(err)
	}
	log.Println("初始化配置成功")
}
