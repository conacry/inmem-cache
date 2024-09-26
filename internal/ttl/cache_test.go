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

func (s *CacheSuite) TestNewCache_WithoutParams_ReturnErr() {
	initParams := CacheInitParam{}

	cache, err := NewCache[string, int](initParams)
	assert.ErrorIs(s.T(), err, ErrIllegalTTL)
	assert.Nil(s.T(), cache)
}

func (s *CacheSuite) TestNewCache_WithAllParams_ReturnCache() {
	initParams := CacheInitParam{
		TTL:      100 * time.Millisecond,
		Capacity: 100,
	}

	cache, err := NewCache[string, int](initParams)
	assert.NotNil(s.T(), cache)
	assert.NoError(s.T(), err)
}

func (s *CacheSuite) TestCache_CacheForInt_StoreInCache() {
	key := "key"
	value := 100500
	ttl := 100 * time.Millisecond

	initParams := CacheInitParam{
		TTL: ttl,
	}
	cache, err := NewCache[string, int](initParams)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), cache)

	err = cache.Set(key, value)
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

	initParams := CacheInitParam{
		TTL: ttl,
	}
	cache, err := NewCache[string, StructForCache](initParams)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), cache)

	err = cache.Set(key, value)
	require.NoError(s.T(), err)

	storedValue, exists := cache.Get(key)
	require.True(s.T(), exists)
	assert.Equal(s.T(), value, storedValue)

	time.Sleep(ttl + 50*time.Millisecond)

	storedValue, exists = cache.Get(key)
	assert.False(s.T(), exists)
	assert.Empty(s.T(), storedValue)
}

func (s *CacheSuite) TestCache_GetByNotExistedKey_ReturnDefaultValue() {
	key := "key"
	value := 100500
	ttl := 100 * time.Millisecond

	initParams := CacheInitParam{
		TTL: ttl,
	}
	cache, err := NewCache[string, int](initParams)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), cache)

	err = cache.Set(key, value)
	require.NoError(s.T(), err)

	storedValue, exists := cache.Get(key)
	require.True(s.T(), exists)
	assert.Equal(s.T(), value, storedValue)

	storedValue, exists = cache.Get("not_existed_key")
	assert.False(s.T(), exists)
	assert.Empty(s.T(), storedValue)
}
