package routes

import (
	"GoBIMS/controllers"
	"GoBIMS/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	// 用户模块路由接口
	router.POST("login/", controllers.Login)
	router.POST("user/add", controllers.AddUser)
	router.GET("users", controllers.GetUser)
	router.PUT("user/", controllers.EditUser)
	router.DELETE("user/:id", controllers.DeleteUser)

	return router
}
