package ttlcache

import (
	"sync"
	"time"
)

type Cache[K comparable, V any] struct {
	data map[K]entry[V]
	ttl  time.Duration
	mu   sync.Mutex
}

func NewCache[K comparable, V any](params CacheInitParam) (*Cache[K, V], error) {
	if params.TTL <= 0 {
		return nil, ErrIllegalTTL
	}

	cache := Cache[K, V]{
		data: make(map[K]entry[V], params.Capacity),
		ttl:  params.TTL,
	}

	return &cache, nil
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

func (c *Cache[K, V]) Set(key K, value V) error {
	item := newEntry[V](value, c.ttl)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = item

	return nil
}

func (c *Cache[K, T]) getZeroValue() T {
	var zeroValue T
	return zeroValue
}
