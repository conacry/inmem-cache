package inmem

import (
	"fmt"
	"time"

	ttlcache "github.com/conacry/inmem-cache/internal/ttl"
	"golang.org/x/exp/constraints"
)

type Cache[K constraints.Ordered, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V, ttl time.Duration) error
}

func NewCache[K constraints.Ordered, V any](cacheType CacheType, opts ...Option) (Cache[K, V], error) {
	switch cacheType {
	case TtlCacheType:
		return makeTtlCache[K, V](opts...)
	case LruCacheType:
		panic("not implemented")
	case LfuCacheType:
		panic("not implemented")
	}

	return nil, fmt.Errorf("unknown cache type: %s", cacheType)
}

func makeTtlCache[K constraints.Ordered, V any](opts ...Option) (Cache[K, V], error) {
	param := CacheInitParam{}
	for _, opt := range opts {
		param = opt(param)
	}

	ttlCacheInitParams := ttlcache.CacheInitParam{
		Size: param.Size,
	}

	return ttlcache.NewCache[K, V](ttlCacheInitParams), nil
}
