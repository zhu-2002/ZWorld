# ZWorld
A homemade social networking site

> @author	Zenos
>
> 行百里者半九十

## 定义用户基本接口

### main.go

```go
package main

import (
	"ZWorld/internal/web"
)

func main() {
	server := web.RegisterRoutes()

	server.Run(":8080")
}
```

### internal/web/init_web.go

```go
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
```

### internal/web/user.go

```go
package web

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户有关的路由
type UserHandler struct {
	emailExp    *regexp2.Regexp
	passwordExp *regexp2.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegexPattern    = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
		passwordRegexPattern = "^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$"
	)
	// 预编译
	emailExp := regexp2.MustCompile(emailRegexPattern, regexp2.None)
	passwordExp := regexp2.MustCompile(passwordRegexPattern, regexp2.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (userHandler *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	// Bind方法会根据 Content-Type 来解析你的数据到 req 中
	// 解析错了，就会直接写回一个 400 错误
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 邮箱
	// 匹配
	ok, err := userHandler.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusBadRequest, "你的邮箱格式不正确")
		return
	}

	// 密码
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusBadRequest, "你的两次密码输入不一致")
		return
	}

	ok, err = userHandler.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusBadRequest, "你的密码格式不正确")
		return
	}

	ctx.String(http.StatusOK, "login success")
	fmt.Println(req)

	// 补充数据库操作

}

func (userHandler *UserHandler) Login(ctx *gin.Context) {

}

func (userHandler *UserHandler) Edit(ctx *gin.Context) {

}

func (userHandler *UserHandler) Profile(ctx *gin.Context) {

}
```

## 跨域问题

请求端和响应端的协议、域名和端口任意一个不同，都是跨域请求。

### middleware

在GIN中提供了一个middleware来解决CORS，[gin-contrib/cors: Official CORS gin's middleware (github.com)](https://github.com/gin-contrib/cors)。

![image-20240122090814819](https://cdn.jsdelivr.net/gh/zhu-2002/img/image-20240122090814819.png)

可以直接使用Gin中Engine上的Use方法来注册middleware，那么进入到这个Engine的所有请求，都会执行对应的代码。

### internal/web/init_web.go

```go
// 路由注册
func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	// 用户路由注册
	RegisterUserRoutes(server)

	// 处理跨域请求
	CORSConfig(server)

	return server
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
```








































## 附录

> 对应源码：
>
> 参考文章链接：
>
> - author：	url：
