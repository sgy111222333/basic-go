package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := InitWebServer()
	//server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, 启动成功")
	})
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
