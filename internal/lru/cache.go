package lrucache

import (
	"fmt"
	"sync"
	"time"
)

type Cache[K comparable, V any] struct {
	data     map[K]Entry[V]
	queue    *AgeList[K]
	capacity int
	mu       sync.Mutex
}

func NewCache[K comparable, V any](params CacheInitParam) (*Cache[K, V], error) {
	if params.Capacity <= 0 {
		return nil, fmt.Errorf("capacity should be greater than 0")
	}

	queue, err := NewAgeList[K](params.Capacity)
	if err != nil {
		return nil, err
	}

	cache := Cache[K, V]{
		data:     make(map[K]Entry[V], params.Capacity),
		queue:    queue,
		capacity: params.Capacity,
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
		c.queue.Remove(key)
		return c.getZeroValue(), false
	}

	return v.value, true
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := NewEntry(value, ttl)
	if _, ok := c.data[key]; ok {
		c.updateEntry(key, item)
	} else {
		c.addNewEntry(key, item)
	}

	return nil
}

func (c *Cache[K, V]) updateEntry(key K, entry Entry[V]) {
	c.queue.MakeYoungest(key)
	c.data[key] = entry
}

func (c *Cache[K, V]) addNewEntry(key K, entry Entry[V]) {
	if len(c.data) >= c.capacity {
		keyToRemove := c.queue.GetOldest()
		c.queue.Remove(keyToRemove)
		delete(c.data, keyToRemove)
	}

	c.data[key] = entry
	c.queue.Add(key)
}

func (c *Cache[K, T]) getZeroValue() T {
	var zeroValue T
	return zeroValue
}
