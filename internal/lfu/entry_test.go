package lfucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntry(t *testing.T) {
	t.Run("Create new entry", func(t *testing.T) {
		value := "value"
		entry := newEntry(value)

		assert.NotEmpty(t, entry)
		assert.Equal(t, value, entry.value)
		assert.Equal(t, 0, entry.useCount)
	})
}
