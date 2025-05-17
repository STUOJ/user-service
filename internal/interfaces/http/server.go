package http

import (
	"stuoj-common/interfaces/http/validator"
	"user-service/pkg/config"

	"github.com/gin-gonic/gin"
)

var (
	ginServer *gin.Engine
)

func InitServer() error {
	config := config.Conf.UserService.Server

	// 创建gin实例
	ginServer = gin.Default()

	validator.SetupValidator()

	// 初始化路由
	err := InitRoute()
	if err != nil {
		return err
	}

	// 启动服务
	err = ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}

	return nil
}
