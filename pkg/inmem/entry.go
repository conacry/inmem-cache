package inmem

type Key string

type CacheEntry[T any] struct {
	Value     T
	ExpiredAt int64
}
