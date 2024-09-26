package inmem

import (
	"time"
)

type CacheInitParam struct {
	Capacity int
	TTL      time.Duration
}

type Option func(param CacheInitParam) CacheInitParam

func WithCapacity(capacity int) Option {
	return func(param CacheInitParam) CacheInitParam {
		param.Capacity = capacity
		return param
	}
}

func WithTTL(ttl time.Duration) Option {
	return func(param CacheInitParam) CacheInitParam {
		param.TTL = ttl
		return param
	}
}
