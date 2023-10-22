package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sgy111222333/basic-go/webook/internal/repository"
	"github.com/sgy111222333/basic-go/webook/internal/repository/dao"
	"github.com/sgy111222333/basic-go/webook/internal/service"
	"github.com/sgy111222333/basic-go/webook/internal/web"
	"github.com/sgy111222333/basic-go/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	initUserHdl(db, server)
	server.Run(":8080")
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}

// middleware 中间件, 其他语言也叫plugin、handler、filter、interceptor
// middleware用来解决一些所有业务都关心的东西, 也叫做AOP(Aspect-Oriented Programming)解决方案

// Gin中接入middleware的方法:
// server.Use(传入一个接收*Context的方法)  // 这个方法就是HandlerFunc

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook?charset=utf8mb4"), &gorm.Config{})
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

func initWebServer() *gin.Engine {
	server := gin.Default()
	// CORS
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type"},
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		//AllowAllOrigins: true, //等价于上面的域名写成"*", 但如果前端引用者策略是strict-origin-when-cross-origin, *不管用
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

	login := &middleware.LoginMiddlewareBuilder{}
	// 存储数据的, 也就是你 userId 存在哪里; 目前直接存cookie里面(不安全)
	store := cookie.NewStore([]byte("secret"))
	//可以User无数次, 每次可以加若干个middleware, 也就是gin.HandlerFunc
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	// sessions.Sessions("ssid", store) 是初始化session
	return server
}
