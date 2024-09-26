package lrucache

import (
	"errors"
)

var (
	ErrIllegalCapacity = errors.New("capacity should be greater than 0")
	ErrIllegalTTL      = errors.New("ttl should be greater than 0")
)
