package main

// Engine
// 一个Web服务器被抽象为Engine, 一个应用可以创建多个Engine实例, 监听不同端口
// Engine承担了路由注册, 接入middleware的核心职责

// Context
// 核心职责是 处理请求 和 返回相应

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	// handler
	server.Use(func(context *gin.Context) {
		println("这是第一个middleware")
	}, func(context *gin.Context) {
		println("这是第二个middleware")
	})

	// 请求方法 + 路由规则 + 处理函数(匿名函数)
	// 静态路由
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, word")
	})
	// 参数路由(路径参数), 如 GET /users/sgy
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello, "+name)
	})
	// 查询参数, 如 GET /order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "订单ID是"+id)
	})
	// 通配符路由, 如 GET /views/index.html  星号不能单独出现, 后面必须有东西
	server.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html")
		// "ABC%s",val   "ABC"+val 这两种形式都可以输出
		ctx.String(http.StatusOK, "view是%s", view) // view是/index.html
	})

	// 如果不传入端口, 默认是8080
	server.Run(":8080") // ! 前面要加冒号
}
