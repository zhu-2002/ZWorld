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
