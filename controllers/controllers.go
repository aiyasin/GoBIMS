package controllers
// 从前端拿到数据
import (
	"GoBIMS/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetBookByID(c *gin.Context) {
	form := c.PostForm("id")
	book := service.GetBookByID(cast.ToInt(form))
	c.IndentedJSON(200, gin.H{
		"message": book,
	})
}

func SearchByName(c *gin.Context) {
	name := c.Query("name")
	fmt.Println(name)
	result := service.SearchByName(name)
	c.IndentedJSON(200, gin.H{
		"data": result,
	})
}
