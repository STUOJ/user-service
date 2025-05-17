package bootstrap

import (
	"log"
	"user-service/internal/interfaces/http"
)

func InitServer() {
	err := http.InitServer()
	if err != nil {
		log.Println("初始化服务器失败！")
		panic(err)
	}
	log.Println("初始化服务器成功")
}
