package cache

import (
	"context"
	"go-cloud-storage/app/user/internal/cache/redis"
	"time"
)

// UserCache 是一个用户信息缓存结构体，它使用 RdbCache 来存储和检索用户数据。
type UserCache struct {
	rdb *redis.RdbCache
}

// NewUserCache 创建并返回一个新的 UserCache 实例。
// 它初始化了 RdbCache，为用户信息缓存提供了一个存储后端。
// NewUserCache 创建用户模块缓存
func NewUserCache() *UserCache {
	return &UserCache{
		rdb: redis.GetRdbCache(),
	}
}

// HSetUserInfo 将用户信息存储到缓存中。
// 它使用哈希表来存储用户信息，允许通过键来唯一标识用户。
// c: 上下文对象，用于传递请求相关的信息。
// key: 在缓存中标识用户的键。
// userInfo: 要存储的用户信息，作为一个接口类型，允许存储任意类型的用户数据。
// 返回错误，如果设置操作失败。
func (uc *UserCache) HSetUserInfo(c context.Context, key string, userInfo interface{}) error {
	if err := uc.rdb.HSet(c, key, userInfo); err != nil {
		return err
	}
	return nil
}

// SetToken 将用户的令牌存储到缓存中，并设置过期时间。
// 这个方法用于存储用户的认证令牌，以便快速检索和验证用户身份。
// c: 上下文对象，用于传递请求相关的信息。
// key: 用于存储令牌的键。
// token: 要存储的用户令牌。
// expire: 令牌的过期时间。
// 返回错误，如果设置操作失败。
func (uc *UserCache) SetToken(c context.Context, key string, token string, expire time.Duration) error {
	if err := uc.rdb.Set(c, key, token, expire); err != nil {
		return err
	}
	return nil
}
