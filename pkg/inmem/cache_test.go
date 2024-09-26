package inmem

import (
	"errors"
	"fmt"
	"testing"
	"time"

	lfucache "github.com/conacry/inmem-cache/internal/lfu"
	lrucache "github.com/conacry/inmem-cache/internal/lru"
	ttlcache "github.com/conacry/inmem-cache/internal/ttl"
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

	cache, err := NewCache[string, string](unknownCacheType)
	assert.Nil(s.T(), cache)
	require.Error(s.T(), err)
	assert.Equal(s.T(), expectedError.Error(), err.Error())
}

func (s *CacheSuite) TestNewCache_TTLCacheTypeWithoutTTL_ReturnError() {
	lruCacheType := TtlCacheType
	cacheErr := errors.New("ttl should be greater than 0")
	expectedErr := fmt.Errorf("failed to create TTL cache: %w", cacheErr)

	cache, err := NewCache[string, string](lruCacheType)
	assert.Nil(s.T(), cache)
	require.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr.Error(), err.Error())
}

func (s *CacheSuite) TestNewCache_TtlCacheType_ReturnCache() {
	ttlCacheType := TtlCacheType
	opts := []Option{
		WithTTL(100 * time.Millisecond),
	}

	cache, err := NewCache[string, string](ttlCacheType, opts...)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), cache)
	assert.IsType(s.T(), &ttlcache.Cache[string, string]{}, cache)
}

func (s *CacheSuite) TestNewCache_LruCacheTypeWithoutCapacity_ReturnError() {
	lruCacheType := CacheType("lru")
	cacheErr := fmt.Errorf("capacity should be greater than 0")
	expectedErr := fmt.Errorf("failed to create LRU cache: %w", cacheErr)

	cache, err := NewCache[string, string](lruCacheType)
	assert.Nil(s.T(), cache)
	require.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr.Error(), err.Error())
}

func (s *CacheSuite) TestNewCache_LruCacheTypeWithoutTTL_ReturnError() {
	lruCacheType := LruCacheType
	opts := []Option{
		WithCapacity(50),
	}
	cacheErr := fmt.Errorf("ttl should be greater than 0")
	expectedErr := fmt.Errorf("failed to create LRU cache: %w", cacheErr)

	cache, err := NewCache[string, string](lruCacheType, opts...)
	assert.Nil(s.T(), cache)
	require.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr.Error(), err.Error())
}

func (s *CacheSuite) TestNewCache_LruCacheType_ReturnCache() {
	lruCacheType := LruCacheType
	opts := []Option{
		WithCapacity(50),
		WithTTL(100 * time.Millisecond),
	}

	cache, err := NewCache[string, string](lruCacheType, opts...)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), cache)
	assert.IsType(s.T(), &lrucache.Cache[string, string]{}, cache)
}

func (s *CacheSuite) TestNewCache_LfuCacheTypeWithoutCapacity_ReturnError() {
	lfuCacheType := LfuCacheType
	cacheErr := fmt.Errorf("capacity should be greater than 0")
	expectedErr := fmt.Errorf("failed to create LFU cache: %w", cacheErr)

	cache, err := NewCache[string, string](lfuCacheType)
	assert.Nil(s.T(), cache)
	require.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr.Error(), err.Error())
}

func (s *CacheSuite) TestNewCache_LfuCacheType_ReturnCache() {
	lfuCacheType := LfuCacheType
	opts := []Option{
		WithCapacity(50),
	}

	cache, err := NewCache[string, string](lfuCacheType, opts...)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), cache)
	assert.IsType(s.T(), &lfucache.Cache[string, string]{}, cache)
}
