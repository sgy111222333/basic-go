//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sgy111222333/basic-go/webook/internal/repository"
	"github.com/sgy111222333/basic-go/webook/internal/repository/cache"
	"github.com/sgy111222333/basic-go/webook/internal/repository/dao"
	"github.com/sgy111222333/basic-go/webook/internal/service"
	"github.com/sgy111222333/basic-go/webook/internal/web"
	"github.com/sgy111222333/basic-go/webook/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		ioc.InitDB, ioc.InitRedis,
		// DAO 部分
		dao.NewUserDAO,
		// Cache 部分
		cache.NewCodeCache, cache.NewUserCache,
		// Repository 部分
		repository.NewCachedUserRepository, repository.NewCodeRepository,
		// Service 部分
		ioc.InitSMSService,
		service.NewUserService,
		service.NewCodeService,
		// Handler 部分
		web.NewUserHandler,
		ioc.InitGinMiddleWares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
