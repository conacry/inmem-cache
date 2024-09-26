package lrucache

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

func (s *CacheSuite) TestNewCache_IllegalCapacity_ReturnError() {
	illegalCapacity := 0

	params := InitParam{Capacity: illegalCapacity}
	cache, err := NewCache[string, struct{}](params)
	assert.Nil(s.T(), cache)
	assert.ErrorIs(s.T(), err, ErrIllegalCapacity)
}

func (s *CacheSuite) TestNewCache_IllegalTTL_ReturnError() {
	correctCapacity := 50
	illegalTTL := 0 * time.Millisecond

	params := InitParam{
		Capacity: correctCapacity,
		TTL:      illegalTTL,
	}
	cache, err := NewCache[string, struct{}](params)
	assert.Nil(s.T(), cache)
	assert.ErrorIs(s.T(), err, ErrIllegalTTL)
}

func (s *CacheSuite) TestNewCache_CorrectInitParams_ReturnCache() {
	correctCapacity := 50
	correctTTL := 100 * time.Millisecond

	params := InitParam{
		Capacity: correctCapacity,
		TTL:      correctTTL,
	}
	cache, err := NewCache[string, struct{}](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)
}

func (s *CacheSuite) TestCache_TtlIsNotExpired_CacheReturnedStoredValue() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := InitParam{
		Capacity: 1,
		TTL:      100 * time.Millisecond,
	}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	key := "key"
	value := StructForCache{
		Field1: "field1",
		Field2: 100500,
	}

	err = cache.Set(key, value)
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

	params := InitParam{
		Capacity: 1,
		TTL:      100 * time.Millisecond,
	}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	key := "key"
	value := StructForCache{
		Field1: "field1",
		Field2: 100500,
	}

	err = cache.Set(key, value)
	require.NoError(s.T(), err)

	time.Sleep(150 * time.Millisecond)

	storedValue, exists := cache.Get(key)
	assert.False(s.T(), exists)
	assert.Empty(s.T(), storedValue)
}

func (s *CacheSuite) TestCache_NotEnoughCapacity_TheOldestValueWasEvicted() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := InitParam{
		Capacity: 1,
		TTL:      100 * time.Millisecond,
	}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	keyOne := "keyOne"
	valueOne := StructForCache{
		Field1: "fieldStructOne1",
		Field2: 100500,
	}

	keyTwo := "keyTwo"
	valueTwo := StructForCache{
		Field1: "fieldStructTwo1",
		Field2: 100501,
	}

	err = cache.Set(keyOne, valueOne)
	require.NoError(s.T(), err)
	assert.Len(s.T(), cache.data, 1)

	err = cache.Set(keyTwo, valueTwo)
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

	params := InitParam{
		Capacity: 1,
		TTL:      100 * time.Millisecond,
	}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	keyOne := "keyOne"
	valueOne := StructForCache{
		Field1: "fieldStructOne1",
		Field2: 100500,
	}

	valueTwo := StructForCache{
		Field1: "fieldStructTwo1",
		Field2: 100501,
	}

	err = cache.Set(keyOne, valueOne)
	require.NoError(s.T(), err)
	assert.Len(s.T(), cache.data, 1)

	storedValueOne, exists := cache.Get(keyOne)
	assert.True(s.T(), exists)
	assert.Equal(s.T(), valueOne, storedValueOne)

	err = cache.Set(keyOne, valueTwo)
	require.NoError(s.T(), err)
	assert.Len(s.T(), cache.data, 1)

	storedValueTwo, exists := cache.Get(keyOne)
	assert.True(s.T(), exists)
	assert.Equal(s.T(), valueTwo, storedValueTwo)
}
