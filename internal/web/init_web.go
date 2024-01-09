package web

import "github.com/gin-gonic/gin"

// 路由注册
func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	// 用户路由注册
	RegisterUserRoutes(server)
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
