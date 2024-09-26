package lfucache

import (
	"time"
)

type bucket[V comparable] map[V]time.Time

type frequencySet[V comparable] struct {
	data     map[int]bucket[V]
	minCount int
}

func newFrequencySet[V comparable]() *frequencySet[V] {
	return &frequencySet[V]{
		data: make(map[int]bucket[V]),
	}
}

func (f *frequencySet[V]) Add(value V) {
	if _, ok := f.data[0]; !ok {
		f.data[0] = make(bucket[V])
	}
	f.data[0][value] = time.Now()
	f.minCount = 0
}

func (f *frequencySet[V]) Touch(value V, count int) {
	delete(f.data[count], value)

	if _, ok := f.data[count+1]; !ok {
		f.data[count+1] = make(bucket[V])
	}

	f.data[count+1][value] = time.Now()

	if count == f.minCount && len(f.data[count]) == 0 {
		f.minCount++
	}
}

func (f *frequencySet[V]) GetLeastFrequent() V {
	minCreatedAt := time.Now()
	var minValue V

	bucket := f.data[f.minCount]
	for value, createdAt := range bucket {
		if createdAt.Before(minCreatedAt) {
			minCreatedAt = createdAt
			minValue = value
		}
	}

	return minValue
}

func (f *frequencySet[V]) Remove(value V, count int) {
	delete(f.data[count], value)
}
