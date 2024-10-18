package ioc

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sgy111222333/basic-go/webook/internal/web"
	"github.com/sgy111222333/basic-go/webook/internal/web/middleware"
	ratelimit "github.com/sgy111222333/basic-go/webook/pkg/ginx/middlerware/ratelimite"
	"strings"
	"time"
)

// InitWebServer mdls是middlewares的简写
func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddleWares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		// CORS
		cors.New(cors.Config{
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
		}),
		// 证明middleware起效
		func(context *gin.Context) {
			fmt.Println("这是第二个middleware")
		},
		// 限流
		ratelimit.NewBuilder(redisClient, time.Second*1, 100).Build(),
		// JWT登录
		(&middleware.LoginJWTMiddlewareBuilder{}).CheckLogin(),
	}
}
