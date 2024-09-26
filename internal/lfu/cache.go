package lfucache

import (
	"sync"
)

type Cache[K comparable, V any] struct {
	data     map[K]*entry[V]
	freq     *frequencySet[K]
	capacity int
	mu       sync.Mutex
}

func NewCache[K comparable, V any](params InitParam) (*Cache[K, V], error) {
	if params.Capacity <= 0 {
		return nil, ErrIllegalCapacity
	}

	cache := Cache[K, V]{
		data:     make(map[K]*entry[V], params.Capacity),
		freq:     newFrequencySet[K](),
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

	c.freq.Touch(key, v.useCount)
	v.useCount++
	return v.value, true
}

func (c *Cache[K, V]) Set(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.data) >= c.capacity {
		c.evictLessUsedEntry()
	}

	if _, ok := c.data[key]; ok {
		c.updateEntry(key, value)
	} else {
		c.addNewEntry(key, value)
	}

	return nil
}

func (c *Cache[K, V]) updateEntry(key K, value V) {
	entry := c.data[key]
	entry.value = value
}

func (c *Cache[K, V]) addNewEntry(key K, value V) {
	entry := newEntry(value)
	c.data[key] = &entry
	c.freq.Add(key)
}

func (c *Cache[K, T]) getZeroValue() T {
	var zeroValue T
	return zeroValue
}

func (c *Cache[K, V]) evictLessUsedEntry() {
	keyToRemove := c.freq.GetLeastFrequent()
	entry := c.data[keyToRemove]
	delete(c.data, keyToRemove)
	c.freq.Remove(keyToRemove, entry.useCount)
}
