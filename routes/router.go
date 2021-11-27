package routes

import (
	"GoBIMS/controllers"
	"GoBIMS/middleware"
	"GoBIMS/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/login/index.html")
	p.AddFromFiles("front", "web/login/ourbooks.html")
	return p
}

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	// 用户模块路由接口
	user := r.Group("api/v1/user")
	user.Use(middleware.JwtToken())
	{
		user.POST("login/", controllers.Login)     //登录
		user.POST("joinup", controllers.JoinUp)    //注册
		user.GET("", controllers.GetUser)          //用户列表
		user.PUT(":id", controllers.EditUser)      //用户信息编辑
		user.DELETE(":id", controllers.DeleteUser) //删除用户
	}
	book := r.Group("api/v1/book")
	{
		book.POST("infomation", controllers.GetBook)
	}

	return r
}
