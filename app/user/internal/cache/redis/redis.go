package redis

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	"go-cloud-storage/pkg/config"
	"go-cloud-storage/pkg/logger"
	"time"
)

var _redisCache *RdbCache

type RdbCache struct {
	rdb *redis.Client
}

func GetRdbCache() *RdbCache {
	return _redisCache
}

func InitRedis(conf config.Redis) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		logger.Errorw("ping redis error", "err", err)
		panic(err)
	}
	_redisCache = &RdbCache{rdb: rdb}
}

func (rc *RdbCache) GetKeys(c context.Context, pattern string) ([]string, error) {
	return rc.rdb.Keys(c, pattern).Result()
}

func (rc *RdbCache) KeyExist(c context.Context, key string) (int64, error) {
	return rc.rdb.Exists(c, key).Result()
}

func (rc *RdbCache) Set(c context.Context, key, value string, expire time.Duration) error {
	err := rc.rdb.Set(c, key, value, expire).Err()
	return err
}

func (rc *RdbCache) Get(c context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(c, key).Result()
	return result, err
}

func (rc *RdbCache) Del(c context.Context, key string) error {
	return rc.rdb.Del(c, key).Err()
}

func (rc *RdbCache) Expire(c context.Context, key string, expireTime time.Duration) (bool, error) {
	// 有目标key返回true，没有目标key，返回false
	return rc.rdb.Expire(c, key, expireTime).Result()
}

func (rc *RdbCache) HSet(c context.Context, key string, value interface{}) error {
	return rc.rdb.HSet(c, key, value).Err()
}

func (rc *RdbCache) HGet(c context.Context, key string, filed string) (string, error) {
	return rc.rdb.HGet(c, key, filed).Result()
}

func (rc *RdbCache) HGetAll(c context.Context, key string) (map[string]string, error) {
	return rc.rdb.HGetAll(c, key).Result()
}

func (rc *RdbCache) IncrHMCount(c context.Context, key, field string, incr int64) (int64, error) {
	return rc.rdb.HIncrBy(c, key, field, incr).Result()
}

func (rc *RdbCache) SAdd(c context.Context, key string, value interface{}) (int64, error) {
	// 返回存入数据的数量
	return rc.rdb.SAdd(c, key, value).Result()
}

func (rc *RdbCache) SGetAll(c context.Context, key string) ([]string, error) {
	return rc.rdb.SMembers(c, key).Result()
}

func (rc *RdbCache) SDel(c context.Context, key string, value interface{}) (int64, error) {
	return rc.rdb.SRem(c, key, value).Result()
}

func (rc *RdbCache) SIsExist(c context.Context, key string, value interface{}) (bool, error) {
	return rc.rdb.SIsMember(c, key, value).Result()
}

func (rc *RdbCache) ZSet(c context.Context, key string, score []float64, member []interface{}) (int64, error) {
	// 返回存入数据的数量
	z := make([]redis.Z, 0)
	for i := 0; i < len(score); i++ {
		z = append(z, redis.Z{Score: score[i], Member: member[i]})
	}
	return rc.rdb.ZAdd(c, key, z...).Result()
}

func (rc *RdbCache) ZGetRangeWithScores(c context.Context, key string, start, stop int64) ([]redis.Z, error) {
	return rc.rdb.ZRangeWithScores(c, key, start, stop).Result()
}

func (rc *RdbCache) ZGetRangeByScoreWithScores(c context.Context, key string, min, max string, offset, count int64) ([]redis.Z, error) {
	return rc.rdb.ZRangeByScoreWithScores(c, key, &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count}).Result()
}

func (rc *RdbCache) ZGetRevRangeByScoreWithScores(c context.Context, key string, min, max string, offset, count int64) ([]redis.Z, error) {
	return rc.rdb.ZRevRangeByScoreWithScores(c, key, &redis.ZRangeBy{Min: min, Max: max, Offset: offset, Count: count}).Result()
}

func (rc *RdbCache) ZDel(c context.Context, key string, value interface{}) (int64, error) {
	return rc.rdb.ZRem(c, key, value).Result()
}

func (rc *RdbCache) LPush(c context.Context, key string, values ...interface{}) (int64, error) {
	return rc.rdb.LPush(c, key, values).Result()
}

func (rc *RdbCache) LPopCount(c context.Context, key string, count int) ([]string, error) {
	return rc.rdb.LPopCount(c, key, count).Result()
}
