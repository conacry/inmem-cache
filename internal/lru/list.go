package lrucache

type ageList[V comparable] struct {
	data []V
}

func newAgeList[V comparable](capacity int) *ageList[V] {
	list := ageList[V]{
		data: make([]V, 0, capacity),
	}

	return &list
}

func (q *ageList[V]) Add(value V) {
	q.data = append(q.data, value)
	return
}

func (q *ageList[V]) GetOldest() V {
	if len(q.data) == 0 {
		var zeroValue V
		return zeroValue
	}

	return q.data[0]
}

func (q *ageList[V]) MakeYoungest(value V) {
	q.Remove(value)
	q.data = append(q.data, value)
}

func (q *ageList[V]) Remove(value V) {
	valueIndex := q.getIndex(value)

	if valueIndex == -1 {
		return
	}

	q.data = append(q.data[:valueIndex], q.data[valueIndex+1:]...)
}

func (q *ageList[V]) getIndex(value V) int {
	valueIndex := -1

	for i, v := range q.data {
		if v == value {
			valueIndex = i
			break
		}
	}

	return valueIndex
}
