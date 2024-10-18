package repository

import (
	"github.com/sgy111222333/basic-go/wire_demo/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(d *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: d,
	}
}
