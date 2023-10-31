package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
}

func (m *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			// 不需要登录校验
			return
		}
		sess := sessions.Default(ctx)
		userID := sess.Get("userID")
		if sess.Get("userId") == nil {
			// 中断, 401, 不执行后面的业务逻辑
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		now := time.Now()
		// 我怎么知道要刷新登录状态的session了呢?
		const updateTimeKey = "update_time"
		val := sess.Get(updateTimeKey)
		lastUpdateTime, ok := val.(time.Time) // 断言val的类型是time.Time
		if val == nil || !ok || now.Sub(lastUpdateTime) > time.Second*10 {
			// 如果第一次进来 || val的类型不对 || 距离上次刷新超过十秒, 把时间存进去
			sess.Set(updateTimeKey, now)
			sess.Set("userId", userID)
			err := sess.Save()
			if err != nil {
				// 打印日志即可, 因为不影响业务
				fmt.Println(err)
			}
		}
	}
}
