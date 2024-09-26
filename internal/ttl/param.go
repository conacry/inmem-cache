package ttlcache

import (
	"time"
)

type CacheInitParam struct {
	Capacity int
	TTL      time.Duration
}
