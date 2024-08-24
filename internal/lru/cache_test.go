package lrucache

import (
	"fmt"
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

func (s *CacheSuite) TestNewCache_IllegalCapacity_ReturnError() {
	illegalCapacity := 0
	expectedErr := fmt.Errorf("capacity should be greater than 0")

	params := CacheInitParam{Capacity: illegalCapacity}
	cache, err := NewCache[string, struct{}](params)
	assert.Nil(s.T(), cache)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr.Error(), err.Error())
}

func (s *CacheSuite) TestNewCache_CorrectCapacity_ReturnCache() {
	correctCapacity := 50

	params := CacheInitParam{Capacity: correctCapacity}
	cache, err := NewCache[string, struct{}](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)
}

func (s *CacheSuite) TestCache_TtlIsNotExpired_CacheReturnedStoredValue() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := CacheInitParam{Capacity: 1}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	key := "key"
	value := StructForCache{
		Field1: "field1",
		Field2: 100500,
	}
	ttl := 10 * time.Millisecond

	err = cache.Set(key, value, ttl)
	require.NoError(s.T(), err)

	storedValue, exists := cache.Get(key)
	assert.True(s.T(), exists)
	assert.Equal(s.T(), value, storedValue)
}

func (s *CacheSuite) TestCache_TtlIsExpired_CacheWasNotReturnStoredValue() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := CacheInitParam{Capacity: 1}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	key := "key"
	value := StructForCache{
		Field1: "field1",
		Field2: 100500,
	}
	ttl := 10 * time.Millisecond

	err = cache.Set(key, value, ttl)
	require.NoError(s.T(), err)

	time.Sleep(50 * time.Millisecond)

	storedValue, exists := cache.Get(key)
	assert.False(s.T(), exists)
	assert.Empty(s.T(), storedValue)
}

func (s *CacheSuite) TestCache_NotEnoughCapacity_TheOldestValueWasEvicted() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := CacheInitParam{Capacity: 1}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	keyOne := "keyOne"
	valueOne := StructForCache{
		Field1: "fieldStructOne1",
		Field2: 100500,
	}
	ttlOne := 50 * time.Millisecond

	keyTwo := "keyTwo"
	valueTwo := StructForCache{
		Field1: "fieldStructTwo1",
		Field2: 100501,
	}
	ttlTwo := 50 * time.Millisecond

	err = cache.Set(keyOne, valueOne, ttlOne)
	require.NoError(s.T(), err)
	assert.Len(s.T(), cache.data, 1)

	err = cache.Set(keyTwo, valueTwo, ttlTwo)
	require.NoError(s.T(), err)
	assert.Len(s.T(), cache.data, 1)

	storedValueOne, exists := cache.Get(keyOne)
	assert.False(s.T(), exists)
	assert.Empty(s.T(), storedValueOne)

	storedValueTwo, exists := cache.Get(keyTwo)
	assert.True(s.T(), exists)
	assert.Equal(s.T(), valueTwo, storedValueTwo)
}

func (s *CacheSuite) TestCache_UsedTheSameKey_ValueInCacheWasUpdated() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := CacheInitParam{Capacity: 1}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	keyOne := "keyOne"
	valueOne := StructForCache{
		Field1: "fieldStructOne1",
		Field2: 100500,
	}
	ttlOne := 50 * time.Millisecond

	valueTwo := StructForCache{
		Field1: "fieldStructTwo1",
		Field2: 100501,
	}
	ttlTwo := 50 * time.Millisecond

	err = cache.Set(keyOne, valueOne, ttlOne)
	require.NoError(s.T(), err)
	assert.Len(s.T(), cache.data, 1)

	storedValueOne, exists := cache.Get(keyOne)
	assert.True(s.T(), exists)
	assert.Equal(s.T(), valueOne, storedValueOne)

	err = cache.Set(keyOne, valueTwo, ttlTwo)
	require.NoError(s.T(), err)
	assert.Len(s.T(), cache.data, 1)

	storedValueTwo, exists := cache.Get(keyOne)
	assert.True(s.T(), exists)
	assert.Equal(s.T(), valueTwo, storedValueTwo)
}
