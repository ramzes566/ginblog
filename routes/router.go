package routes

import (
	"ginblog/utils"

	v1 "ginblog/api/v1"

	"github.com/gin-contrib/multitemplate"

	"ginblog/middleware"

	"github.com/gin-gonic/gin"
)

func createrMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	// p.AddFromFiles("admin", "web/admin/dist/index.html")
	// p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createrMyRender()
	r.Use(middleware.Logger()) // 日志中间件
	r.Use(gin.Recovery())      // recovery中间件
	r.Use(middleware.Cors())   // 跨域中间件

	// r.Static("/static", "./web/front/dist/static")               // 静态文件目录
	// r.Static("/admin", "./web/admin/dist")                       // 静态文件目录
	// r.StaticFile("/favicon.ico", "./web/front/dist/favicon.ico") // favicon.ico

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken()) // jwt中间件
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)   // 查询用户列表
		auth.PUT("user/:id", v1.EditUser)      // 编辑用户信息
		auth.DELETE("user/:id", v1.DeleteUser) // 删除用户
	}

	router := r.Group("/api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)    // 添加用户
		router.GET("user/:id", v1.GetUserInfo) // 查询单个用户信息
		router.GET("users", v1.GetUsers)       // 查询用户列表
	}
	_ = r.Run(utils.HttpPort)
}
