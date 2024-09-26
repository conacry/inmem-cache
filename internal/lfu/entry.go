package lfucache

type entry[T any] struct {
	value    T
	useCount int
}

func newEntry[T any](value T) entry[T] {
	return entry[T]{
		value:    value,
		useCount: 0,
	}
}
