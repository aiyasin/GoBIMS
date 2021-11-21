package routes

import (
	"GoBIMS/controllers"
	"GoBIMS/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.POST("/api/v1/book", controllers.GetBookByID)
	router.GET("/api/v1/search", controllers.SearchByName)
	return router
}
