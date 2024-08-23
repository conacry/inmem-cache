package ttlcache

import (
	"testing"
	"time"

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

func (s *CacheSuite) TestCache_NewCache_WithoutParams() {
	initParams := CacheInitParam{}

	cache := NewCache[string, int](initParams)
	assert.NotNil(s.T(), cache)
}

func (s *CacheSuite) TestCache_NewCache_WithParams() {
	initParams := CacheInitParam{
		Size: 100,
	}

	cache := NewCache[string, int](initParams)
	assert.NotNil(s.T(), cache)
}

func (s *CacheSuite) TestCache_CacheForInt_StoreInCache() {
	key := "key"
	value := 100500
	ttl := 100 * time.Millisecond

	initParams := CacheInitParam{}
	cache := NewCache[string, int](initParams)

	err := cache.Set(key, value, ttl)
	require.NoError(s.T(), err)

	storedValue, exists := cache.Get(key)
	require.True(s.T(), exists)
	assert.Equal(s.T(), value, storedValue)

	time.Sleep(ttl + 50*time.Millisecond)

	storedValue, exists = cache.Get(key)
	assert.False(s.T(), exists)
	assert.Empty(s.T(), storedValue)
}

func (s *CacheSuite) TestCache_CacheForStruct_StoreInCache() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	key := "key"
	value := StructForCache{
		Field1: "field_value",
		Field2: 100500,
	}
	ttl := 100 * time.Millisecond

	initParams := CacheInitParam{}
	cache := NewCache[string, StructForCache](initParams)

	err := cache.Set(key, value, ttl)
	require.NoError(s.T(), err)

	storedValue, exists := cache.Get(key)
	require.True(s.T(), exists)
	assert.Equal(s.T(), value, storedValue)

	time.Sleep(ttl + 50*time.Millisecond)

	storedValue, exists = cache.Get(key)
	assert.False(s.T(), exists)
	assert.Empty(s.T(), storedValue)
}
