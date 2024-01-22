package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// 路由注册
func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	// 用户路由注册
	RegisterUserRoutes(server)

	// 处理跨域请求
	CORSConfig(server)

	return server
}

// 用户路由注册
func RegisterUserRoutes(server *gin.Engine) {
	userHandler := &UserHandler{}
	userHandler = NewUserHandler()
	// 集中路由
	ug := server.Group("/users")
	ug.POST("/signup", userHandler.SignUp)
	ug.POST("/login", userHandler.Login)
	ug.POST("/edit", userHandler.Edit)
	ug.GET("/profile", userHandler.Profile)
}

// 解决跨域问题
func CORSConfig(server *gin.Engine) {
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	server.Use(cors.New(cors.Config{
		// AllowOrigins: []string{"http://localhost:3000"},
		// AllowMethods:     []string{},
		AllowHeaders: []string{"Authorization", "content-type"},
		// 是否允许携带cookie
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// 开发环境
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yoursever.com")
		},
		MaxAge: 12 * time.Hour,
	}))
}
