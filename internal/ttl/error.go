package ttlcache

import (
	"errors"
)

var (
	ErrIllegalTTL = errors.New("ttl should be greater than 0")
)
