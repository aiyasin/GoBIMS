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
	p.AddFromFiles("front", "web/html/login.html")
	return p
}

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Static("/static", "./web/html/")
	r.StaticFile("/favicon.svg", "./web/favicon.svg")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})
	// 基础模块路由接口
	base := r.Group("api/v1/base")
	base.Use(middleware.JwtToken()) //使用JWT中间件做token接口
	{
		base.POST("login", controllers.LoginFront)     //前台登录
		base.POST("joinup", controllers.JoinUp)        //注册
		base.GET("booklist", controllers.GetBookList)  //图书信息列表
		base.GET("searchlist", controllers.SearchBook) //搜索图书
	}
	// 用户模块路由接口
	user := r.Group("api/v1/user")
	// user.Use(middleware.JwtToken())
	{
		user.PUT(":id", controllers.EditUser)      //编辑用户信息
		user.DELETE(":id", controllers.DeleteUser) //删除用户
	}

	// 管理员模块路由接口
	admin := r.Group("api/v1/admin")
	{

		admin.GET("userlist", controllers.GetUser)             //用户列表or使用username搜索
		admin.GET("searchuser", controllers.GetUser)           //用户查询
		admin.POST("login", controllers.Login)                 //后台登陆
		admin.POST("addbook", controllers.AddBook)             //增加图书
		admin.DELETE("deletebook/:id", controllers.DeleteBook) //删除图书
		admin.PUT("editbook/:id", controllers.EditBook)        //编辑图书信息
	}
	return r
}
