package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sgy111222333/basic-go/webook/internal/domain"
	"time"
)

var ErrKeyNotExist = redis.Nil

type UserCache struct {
	// ! 这里用的是redis.Cmdable这个接口, 而不是redis.Client, 因为要面向接口编程
	cmd        redis.Cmdable
	expiration time.Duration
}

func (c *UserCache) Set(ctx context.Context, du domain.User) error {
	key := c.key(du.Id)
	data, err := json.Marshal(du)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()
}

func (c *UserCache) Get(ctx context.Context, uid int64) (domain.User, error) {
	key := c.key(uid)
	data, err := c.cmd.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	// Set时用 JSON 来序列化, 所以这里用 JSON 反序列化
	err = json.Unmarshal([]byte(data), &u)
	return domain.User{}, err
}

func (c *UserCache) key(uid int64) string {
	// user-info	user.info	user|info	user_info	用常见的分隔符都可以
	return fmt.Sprintf("user:info:%d", uid)
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}
