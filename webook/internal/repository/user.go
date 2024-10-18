package repository

import (
	"context"
	"database/sql"
	"github.com/sgy111222333/basic-go/webook/internal/domain"
	"github.com/sgy111222333/basic-go/webook/internal/repository/cache"
	"github.com/sgy111222333/basic-go/webook/internal/repository/dao"
	"log"
	"time"
)

var (
	ErrDuplicateUser = dao.ErrDuplicateEmail
	ErrUserNotFound  = dao.ErrRecordNotFound
)

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateNonZeroFields(ctx context.Context, user domain.User) error
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	FindById(ctx context.Context, uid int64) (domain.User, error)
}

type CachedUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

type DBConfig struct {
	DSN string
}

type CacheConfig struct {
	Addr string
}

func NewCachedUserRepository(dao dao.UserDAO, cache cache.UserCache) UserRepository {
	return &CachedUserRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *CachedUserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, repo.toEntity(u))
	//return repo.dao.Insert(ctx, dao.User{
	//	Email:    sql.NullString{},
	//	Password: u.Password,
	//})
}

func (repo *CachedUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *CachedUserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email.String,
		Phone:    u.Phone.String,
		Password: u.Password,
		AboutMe:  u.AboutMe,
		Nickname: u.Nickname,
		Birthday: time.UnixMilli(u.Birthday),
	}
}

func (repo *CachedUserRepository) toEntity(u domain.User) dao.User {
	return dao.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Password: u.Password,
		AboutMe:  u.AboutMe,
		Nickname: u.Nickname,
		Birthday: u.Birthday.UnixMilli(),
	}
}
func (repo *CachedUserRepository) UpdateNonZeroFields(ctx context.Context, user domain.User) error {
	return repo.dao.UpdateById(ctx, repo.toEntity(user))
}

// FindById 普通场景, redis挂了之后查DB
func (repo *CachedUserRepository) FindById(ctx context.Context, uid int64) (domain.User, error) {
	// 先查redis, du是domain user的意思
	du, err := repo.cache.Get(ctx, uid)
	// 如果查缓存不报错, 说明拿到了想要的数据, 直接返回
	if err == nil {
		return du, nil
	}
	// 只有查缓存报错(err不为空), 才会走到查DB这步
	// err的两种可能: 1. key不存在, redis正常; 2. redis连不上
	u, err := repo.dao.FindById(ctx, uid)
	if err != nil {
		return domain.User{}, err
	}
	du = repo.toDomain(u)
	// 查完DB后把数据存入redis, 下一次请求就快了
	// 存redis可以异步执行
	go func() {
		err = repo.cache.Set(ctx, du)
		if err != nil {
			// 写不进去可能是redis/网络挂了, 下次查询还是查DB, 这就是缓存击穿
			log.Println(err)
		}
	}()

	//err = repo.cache.Set(ctx, du)
	//if err != nil {
	//	log.Println(err)
	//}

	return du, nil
}

// FindByIdV1 高并发场景, redis挂之后业务不可用
func (repo *CachedUserRepository) FindByIdV1(ctx context.Context, uid int64) (domain.User, error) {
	// 先查redis, du是domain user的意思
	du, err := repo.cache.Get(ctx, uid)
	// 如果查缓存不报错, 说明拿到了想要的数据, 直接返回
	switch err {
	case nil:
		return du, nil
	case cache.ErrKeyNotExist:
		// key不存在, 从BD查, 并存入缓存
		u, err := repo.dao.FindById(ctx, uid)
		if err != nil {
			return domain.User{}, err
		}
		du = repo.toDomain(u)
		err = repo.cache.Set(ctx, du)
		if err != nil {
			log.Println(err)
		}
		return du, nil
	default:
		// 高并发用降级的写法, redis不正常, DB也不去查了, 防止打垮DB, 让此业务不可用, 因为万一DB挂了, 可能影响一堆业务
		return domain.User{}, err
	}
}

func (repo *CachedUserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := repo.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, nil
	}
	return repo.toDomain(u), err
}
