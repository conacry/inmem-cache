package lfucache

import (
	"errors"
)

var (
	ErrIllegalCapacity = errors.New("capacity should be greater than 0")
)
