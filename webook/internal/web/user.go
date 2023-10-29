package web

import (
	//"regexp" // go官方的regexp不支持?=.的写法
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sgy111222333/basic-go/webook/internal/domain"
	"github.com/sgy111222333/basic-go/webook/internal/service"
	"net/http"
	"time"
)

const (
	// 也可以用``括起来, 这样没有转译
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,16}$`
)

// UserHandler 所有和用户有关的路由都定义在这里
type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserService
}

// NewUserHandler 预编译正则表达式
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:            svc,
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

/*
	集中注册: 在main就能看到所有路由, 且在git里容易冲突
	分散注册: 一类handler一个文件, 更有条理
*/
func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	// REST 风格
	//server.POST("user",h.SignUP())
	//server.POST("/user", h.SignUP())
	//server.GET("/user/:id", h.SignUP())
	// 未使用分组路由
	//server.POST("/users/signup", h.SignUP)
	//server.POST("/users/login", h.Login)
	//server.POST("/users/edit", h.Edit)
	//server.GET("/users/profile", h.Profile)
	// 分组路由
	ug := server.Group("/users") // 把users拼在前面
	ug.POST("/signup", h.SignUP)
	//ug.POST("/login", h.Login)
	ug.POST("/login", h.LoginJWT)
	ug.POST("/edit", h.Edit)
	ug.GET("/profile", h.Profile)
}

func (h *UserHandler) SignUP(ctx *gin.Context) {
	// 也可以把结构体放到方法外面, 但不是最小范围了
	type SignUpReq struct {
		Email           string `json:"email"` // 这个叫标签
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq
	// Bind 把前端传过来的json 根据标签 填入结构体; 格式不对会返回400
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 不使用预编译正则表达式
	//isEmail, err := regexp.MatchString(emailRegexPattern, req.Email)
	// 使用预编译正则表达式
	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil { // 匹配超时
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱格式不对")
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入密码不一致")
		return
	}
	isPassword, err := h.passwordRexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码格式不对, 必须包含大小写字母和数字的组合，可以使用特殊字符，长度在8-16之间")
		return
	}

	err = h.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	// 要判定邮箱冲突: 拿到数据库的唯一索引冲突错误
	switch err {
	case nil:
		ctx.String(http.StatusOK, "注册成功")
	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "该邮箱已被注册")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}
func (h *UserHandler) LoginJWT(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		uc :=
			UserClaims{
				Uid: u.Id,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)), // 1分钟登录 过期
				},
				UserAgent: ctx.GetHeader("User-Agent"),
			}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
		tokenStr, err := token.SignedString(JWTKey)
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
		}
		ctx.Header("x-jwt-token", tokenStr)
		ctx.String(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或密码错误")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}
func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		sess := sessions.Default(ctx)
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			MaxAge: 900, // 十五分钟
			//HttpOnly: true,
		})
		err = sess.Save()
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
			return
		}
		ctx.String(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或密码错误")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}
func (h *UserHandler) Edit(ctx *gin.Context) {

}
func (h *UserHandler) Profile(ctx *gin.Context) {
	//uc := ctx.MustGet("user").(UserClaims)
	ctx.String(http.StatusOK, "这是 profile")
}

var JWTKey = []byte("7tY76KDM3z6P2jWykvNt7eRaX7AYqRmR")

type UserClaims struct {
	jwt.RegisteredClaims // 组合
	Uid                  int64
	UserAgent            string
}
