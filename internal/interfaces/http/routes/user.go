package routes

import (
	"user-service/internal/interfaces/http/handler"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(ginServer *gin.Engine) {
	userRoute := ginServer.Group("/user")
	{
		userRoute.GET("/:id", handler.UserInfo)
		userRoute.POST("/login", handler.UserLogin)
		userRoute.POST("/register", handler.UserRegister)
		userRoute.PUT("/password", handler.UserChangePassword)

		userRoute.GET("/current", handler.UserCurrentId)
		userRoute.PUT("/modify", handler.UserModify)
		userRoute.POST("/avatar", handler.ModifyUserAvatar)

		userRoute.GET("/", handler.UserList)
		userRoute.POST("/", handler.UserAdd)
		userRoute.PUT("/role", handler.UserModifyRole)
		userRoute.GET("/statistics", handler.UserStatistics)
	}
}
