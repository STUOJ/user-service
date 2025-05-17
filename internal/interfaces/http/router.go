package http

import (
	"net/http"
	"stuoj-common/interfaces/http/vo"
	"user-service/internal/interfaces/http/routes"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, vo.RespOk("user-service 启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, vo.RespError("404 Not Found", nil))
	})

	// 初始化路由
	routes.InitUserRoute(ginServer)

	return nil
}
