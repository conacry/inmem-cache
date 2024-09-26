package lrucache

import (
	"time"
)

type InitParam struct {
	Capacity int
	TTL      time.Duration
}
