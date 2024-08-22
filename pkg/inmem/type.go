package inmem

type CacheType string

const (
	TtlCacheType CacheType = "ttl"
	LruCacheType CacheType = "lru"
	LfuCacheType CacheType = "lfu"
)
