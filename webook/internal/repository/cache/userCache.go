package cache

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache struct {
	// ! 这里用的是redis.Cmdable这个接口, 而不是redis.Client, 因为要面向接口编程
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}
