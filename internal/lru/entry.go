package lrucache

import (
	"time"
)

type Entry[T any] struct {
	value     T
	expiredAt time.Time
}

func NewEntry[T any](value T, ttl time.Duration) Entry[T] {
	return Entry[T]{
		value:     value,
		expiredAt: time.Now().Add(ttl),
	}
}
