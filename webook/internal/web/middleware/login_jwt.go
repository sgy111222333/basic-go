package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sgy111222333/basic-go/webook/internal/web"
	"log"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
}

func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			return
		}
		// 根据约定, 从前端拿到的token在Authorization里
		// Bearer xxxxxx.xxxxxx.xxxxxx 的形式, 需要切割
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			// 没登录, 没token
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(authCode, " ")
		if len(segs) != 2 {
			// 切割结果不对, 说明Authorization有问题
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil // 这里返回的是固定的key, 也可以根据路径等内容计算一个动态的key
		})
		if err != nil {
			// token无效
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			// token非法或者过期
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expireTime, err := uc.GetExpirationTime()
		//if expireTime.Before(time.Now()){ // 这样判定也可以
		//	//token过期
		//	ctx.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		// 剩余过期时间小于20s就要刷新
		if expireTime.Sub(time.Now()) < time.Second*20 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 1)) // 增加1分钟的过期时间
			newToken, err := token.SignedString(web.JWTKey)
			if err != nil {
				log.Println(err)
			}
			ctx.Header("x-jwt-token", newToken)
		}
		//ctx.Set("user", uc)
	}
}
