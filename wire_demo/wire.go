//go:build wireinject

// 上面表示让wire的命令行工具来处理这个文件
package wire_demo

import (
	"github.com/google/wire"
	"github.com/sgy111222333/basic-go/wire_demo/repository"
	"github.com/sgy111222333/basic-go/wire_demo/repository/dao"
)

func InitUserRepository() *repository.UserRepository {
	// 函数带括号是发起调用, 函数不带括号是传入函数本身
	wire.Build(repository.NewUserRepository, dao.NewUserDAO, InitDB)
	return &repository.UserRepository{}
}
