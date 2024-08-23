package ttlcache

import (
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

type Cache[K constraints.Ordered, V any] struct {
	data map[K]CacheEntry[V]
	mu   sync.Mutex
}

func NewCache[K constraints.Ordered, V any](initParams CacheInitParam) *Cache[K, V] {
	cache := &Cache[K, V]{
		data: make(map[K]CacheEntry[V], initParams.Size),
	}

	return cache
}

func (c *Cache[K, T]) Get(key K) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.data[key]
	if !ok {
		return c.getZeroValue(), false
	}

	if v.ExpiredAt > 0 && time.Now().UnixMilli() > v.ExpiredAt {
		delete(c.data, key)
		return c.getZeroValue(), false
	}

	return v.Value, true
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) error {
	expiredAt := time.Now().Add(ttl).UnixMilli()

	item := CacheEntry[V]{
		Value:     value,
		ExpiredAt: expiredAt,
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = item

	return nil
}

func (c *Cache[K, T]) getZeroValue() T {
	var zeroValue T
	return zeroValue
}
