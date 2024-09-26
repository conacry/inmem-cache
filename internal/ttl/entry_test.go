package ttlcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEntry(t *testing.T) {
	t.Run("Create new entry", func(t *testing.T) {
		value := "value"
		ttl := 10 * time.Second
		entry := newEntry(value, ttl)

		assert.NotEmpty(t, entry)
		assert.Equal(t, value, entry.value)
		assert.NotEqual(t, time.Time{}, entry.expiredAt)
	})
}
