package inmem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CacheSuite struct {
	suite.Suite
}

func TestCacheSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CacheSuite))
}

func (s *CacheSuite) TestNewCache_UnknownCacheType_ReturnError() {
	unknownCacheType := CacheType("unknown")
	expectedError := fmt.Errorf("unknown cache type: %s", unknownCacheType)

	cache, err := NewCache[string](unknownCacheType)
	assert.Nil(s.T(), cache)
	require.Error(s.T(), err)
	assert.Equal(s.T(), err.Error(), expectedError.Error())
}

func (s *CacheSuite) TestNewCache_TtlCacheType_PanicCacheNotImplemented() {
	ttlCacheType := CacheType("ttl")

	assert.PanicsWithValue(s.T(), "not implemented", func() {
		_, _ = NewCache[string](ttlCacheType)
	})
}

func (s *CacheSuite) TestNewCache_LruCacheType_PanicCacheNotImplemented() {
	lruCacheType := CacheType("lru")

	assert.PanicsWithValue(s.T(), "not implemented", func() {
		_, _ = NewCache[string](lruCacheType)
	})
}

func (s *CacheSuite) TestNewCache_LfuCacheType_PanicCacheNotImplemented() {
	lfuCacheType := CacheType("lfu")

	assert.PanicsWithValue(s.T(), "not implemented", func() {
		_, _ = NewCache[string](lfuCacheType)
	})
}
