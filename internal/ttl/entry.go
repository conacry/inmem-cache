package ttlcache

type CacheEntry[T any] struct {
	Value     T
	ExpiredAt int64 // UnixMilli
}
