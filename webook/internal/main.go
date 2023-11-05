package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sgy111222333/basic-go/webook/config"
	"github.com/sgy111222333/basic-go/webook/internal/repository"
	"github.com/sgy111222333/basic-go/webook/internal/repository/cache"
	"github.com/sgy111222333/basic-go/webook/internal/repository/dao"
	"github.com/sgy111222333/basic-go/webook/internal/service"
	"github.com/sgy111222333/basic-go/webook/internal/web"
	"github.com/sgy111222333/basic-go/webook/internal/web/middleware"
	ratelimit "github.com/sgy111222333/basic-go/webook/pkg/ginx/middlerware/ratelimite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {
	db := initDB()
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	codeSvc := initCodeSvc(redisClient)
	server := initWebServer()
	initUserHdl(db, redisClient, codeSvc, server)
	//server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, 启动成功")
	})
	server.Run(":8080")
}

func initUserHdl(db *gorm.DB, redisClient redis.Cmdable, codeSvc *service.CodeService, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	uc := cache.NewUserCache(redisClient)
	ur := repository.NewUserRepository(ud, uc)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us, codeSvc)
	hdl.RegisterRoutes(server)
}

// middleware 中间件, 其他语言也叫plugin、handler、filter、interceptor
// middleware用来解决一些所有业务都关心的东西, 也叫做AOP(Aspect-Oriented Programming)解决方案

// Gin中接入middleware的方法:
// server.Use(传入一个接收*Context的方法)  // 这个方法就是HandlerFunc

func initDB() *gorm.DB {
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook?charset=utf8mb4"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug() // 打印出所有生成的sql语句
	err = dao.InitTables(db)
	if err != nil {
		return nil
	}
	return db
}
func initCodeSvc(redisClient redis.Cmdable) *service.CodeService {
	cc := cache.NewCodeCache(redisClient)
	crepo := repository.NewCodeRepository(cc)
	return service.NewCodeService(crepo, nil)
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	// CORS
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://192.168.6.155:3000"},
		ExposeHeaders:    []string{"x-jwt-token"}, // 允许前端访问我的某个header
		//AllowAllOrigins:  true,                    //等价于上面的域名写成"*", 但如果前端引用者策略是strict-origin-when-cross-origin, *不管用
		AllowOriginFunc: func(origin string) bool {
			//	允许包含localhost和youzu的域名访问
			if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "youzu")
		},
		MaxAge: 12 * time.Hour,
	}), func(context *gin.Context) {
		fmt.Println("这是第二个middleware")
	})
	// * 限流插件, 压测时关闭或调大rate
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second*1, 100).Build())

	useJWT(server)
	//useSession(server)
	return server
}
func useJWT(server *gin.Engine) {
	login := &middleware.LoginJWTMiddlewareBuilder{}
	server.Use(login.CheckLogin())

}
func useSession(server *gin.Engine) {
	login := &middleware.LoginMiddlewareBuilder{}
	// 存储数据的, 也就是你 userId 存在哪里; 目前直接存cookie里面(不安全)
	store := cookie.NewStore([]byte("secret"))
	//store = memstore.NewStore([]byte("jXhyXxLsp2fqQ2TDz42RKMfpkvtJYEQd"), []byte("RbQ3vwLWk6HHRK5auSL8m6TfdsJkgWJz"))
	//store, err := redis.NewStore(16, "tcp", "127.0.0.1:16379", "",
	//	[]byte("jXhyXxLsp2fqQ2TDz42RKMfpkvtJYEQd"), // Authentication 身份认证
	//	[]byte("RbQ3vwLWk6HHRK5auSL8m6TfdsJkgWJz")) // Encryption 数据加密
	//if err != nil {
	//	panic(err)
	//}
	//可以User无数次, 每次可以加若干个middleware, 也就是gin.HandlerFunc
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	// sessions.Sessions("ssid", store) 是初始化session, store是存session的地方, 可以是cookie、memstore、redis、mysql等
	// 可以传入不同store的原因是: Session传入的是interface, 这就是面向接口编程的好处
}
