package inmem

import (
	"fmt"
	"time"
)

type Cache[T any] interface {
	Get(key Key) (T, bool)
	Set(key Key, value T, ttl time.Duration)
}

func NewCache[T any](cacheType CacheType, opts ...Option) (Cache[T], error) {
	switch cacheType {
	case TtlCacheType:
		panic("not implemented")
	case LruCacheType:
		panic("not implemented")
	case LfuCacheType:
		panic("not implemented")
	}

	return nil, fmt.Errorf("unknown cache type: %s", cacheType)
}
