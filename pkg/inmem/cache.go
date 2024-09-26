package inmem

import (
	"fmt"

	lfucache "github.com/conacry/inmem-cache/internal/lfu"
	lrucache "github.com/conacry/inmem-cache/internal/lru"
	ttlcache "github.com/conacry/inmem-cache/internal/ttl"
)

type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V) error
}

func NewCache[K comparable, V any](cacheType CacheType, opts ...Option) (Cache[K, V], error) {
	switch cacheType {
	case TtlCacheType:
		return makeTtlCache[K, V](opts...)
	case LruCacheType:
		return makeLruCache[K, V](opts...)
	case LfuCacheType:
		return makeLfuCache[K, V](opts...)
	default:
		return nil, fmt.Errorf("unknown cache type: %s", cacheType)
	}
}

func makeTtlCache[K comparable, V any](opts ...Option) (Cache[K, V], error) {
	param := CacheInitParam{}
	for _, opt := range opts {
		param = opt(param)
	}

	ttlCacheInitParams := ttlcache.CacheInitParam{
		Capacity: param.Capacity,
		TTL:      param.TTL,
	}

	cache, err := ttlcache.NewCache[K, V](ttlCacheInitParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create TTL cache: %w", err)
	}

	return cache, nil
}

func makeLruCache[K comparable, V any](opts ...Option) (Cache[K, V], error) {
	param := CacheInitParam{}
	for _, opt := range opts {
		param = opt(param)
	}

	lruCacheInitParams := lrucache.InitParam{
		Capacity: param.Capacity,
		TTL:      param.TTL,
	}

	cache, err := lrucache.NewCache[K, V](lruCacheInitParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create LRU cache: %w", err)
	}

	return cache, nil
}

func makeLfuCache[K comparable, V any](opts ...Option) (Cache[K, V], error) {
	param := CacheInitParam{}
	for _, opt := range opts {
		param = opt(param)
	}

	lfuCacheInitParams := lfucache.InitParam{
		Capacity: param.Capacity,
	}

	cache, err := lfucache.NewCache[K, V](lfuCacheInitParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create LFU cache: %w", err)
	}

	return cache, nil
}
