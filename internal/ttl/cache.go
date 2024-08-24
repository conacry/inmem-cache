package ttlcache

import (
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

type Cache[K constraints.Ordered, V any] struct {
	data map[K]Entry[V]
	mu   sync.Mutex
}

func NewCache[K constraints.Ordered, V any](initParams CacheInitParam) *Cache[K, V] {
	cache := Cache[K, V]{
		data: make(map[K]Entry[V], initParams.Size),
	}

	return &cache
}

func (c *Cache[K, T]) Get(key K) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.data[key]
	if !ok {
		return c.getZeroValue(), false
	}

	if !v.expiredAt.IsZero() && time.Now().After(v.expiredAt) {
		delete(c.data, key)
		return c.getZeroValue(), false
	}

	return v.value, true
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) error {
	item := NewEntry[V](value, ttl)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = item

	return nil
}

func (c *Cache[K, T]) getZeroValue() T {
	var zeroValue T
	return zeroValue
}
