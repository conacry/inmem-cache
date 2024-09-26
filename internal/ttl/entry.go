package ttlcache

import (
	"time"
)

type entry[T any] struct {
	value     T
	expiredAt time.Time
}

func newEntry[T any](value T, ttl time.Duration) entry[T] {
	return entry[T]{
		value:     value,
		expiredAt: time.Now().Add(ttl),
	}
}
