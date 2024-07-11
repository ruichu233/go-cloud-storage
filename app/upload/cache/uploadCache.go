package cache

import (
	"context"
)

type UploadCache struct {
	rdb *cache.RdbCache
}

func NewUploadCache() *UploadCache {
	return &UploadCache{
		rdb: cache.GetRdbCache(),
	}
}

func (uc *UploadCache) HSetMultipartUploadInfo(c context.Context, key string, value map[string]interface{}) error {
	if err := uc.rdb.HSet(c, key, value); err != nil {
		return err
	}
	return nil
}
