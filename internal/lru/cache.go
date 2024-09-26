package lrucache

import (
	"sync"
	"time"
)

type Cache[K comparable, V any] struct {
	data     map[K]entry[V]
	list     *ageList[K]
	capacity int
	ttl      time.Duration
	mu       sync.Mutex
}

func NewCache[K comparable, V any](params InitParam) (*Cache[K, V], error) {
	if params.Capacity <= 0 {
		return nil, ErrIllegalCapacity
	}

	if params.TTL <= 0 {
		return nil, ErrIllegalTTL
	}

	cache := Cache[K, V]{
		data:     make(map[K]entry[V], params.Capacity),
		list:     newAgeList[K](params.Capacity),
		capacity: params.Capacity,
		ttl:      params.TTL,
	}

	return &cache, nil
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.data[key]
	if !ok {
		return c.getZeroValue(), false
	}

	if !v.expiredAt.IsZero() && time.Now().After(v.expiredAt) {
		delete(c.data, key)
		c.list.Remove(key)
		return c.getZeroValue(), false
	}

	c.list.MakeYoungest(key)
	return v.value, true
}

func (c *Cache[K, V]) Set(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := newEntry(value, c.ttl)
	if _, ok := c.data[key]; ok {
		c.updateEntry(key, item)
	} else {
		c.addNewEntry(key, item)
	}

	return nil
}

func (c *Cache[K, V]) updateEntry(key K, entry entry[V]) {
	c.list.MakeYoungest(key)
	c.data[key] = entry
}

func (c *Cache[K, V]) addNewEntry(key K, entry entry[V]) {
	if len(c.data) >= c.capacity {
		keyToRemove := c.list.GetOldest()
		c.list.Remove(keyToRemove)
		delete(c.data, keyToRemove)
	}

	c.data[key] = entry
	c.list.Add(key)
}

func (c *Cache[K, T]) getZeroValue() T {
	var zeroValue T
	return zeroValue
}
