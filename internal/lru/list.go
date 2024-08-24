package lrucache

import (
	"fmt"
)

type AgeList[V comparable] struct {
	data []V
}

func NewAgeList[V comparable](capacity int) (*AgeList[V], error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity should be greater than 0")
	}

	list := AgeList[V]{
		data: make([]V, 0, capacity),
	}

	return &list, nil
}

func (q *AgeList[V]) Add(value V) {
	q.data = append(q.data, value)
	return
}

func (q *AgeList[V]) GetOldest() V {
	if len(q.data) == 0 {
		var zeroValue V
		return zeroValue
	}

	return q.data[0]
}

func (q *AgeList[V]) MakeYoungest(value V) {
	q.Remove(value)
	q.data = append(q.data, value)
}

func (q *AgeList[V]) Remove(value V) {
	valueIndex := q.getIndex(value)

	if valueIndex == -1 {
		return
	}

	q.data = append(q.data[:valueIndex], q.data[valueIndex+1:]...)
}

func (q *AgeList[V]) getIndex(value V) int {
	valueIndex := -1

	for i, v := range q.data {
		if v == value {
			valueIndex = i
			break
		}
	}

	return valueIndex
}
