package lfucache

import (
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

func (s *CacheSuite) TestNewCache_IllegalCapacity_ReturnError() {
	illegalCapacity := 0

	params := InitParam{Capacity: illegalCapacity}
	cache, err := NewCache[string, struct{}](params)
	assert.Nil(s.T(), cache)
	assert.ErrorIs(s.T(), err, ErrIllegalCapacity)
}

func (s *CacheSuite) TestNewCache_CorrectCapacity_ReturnCache() {
	correctCapacity := 50

	params := InitParam{Capacity: correctCapacity}
	cache, err := NewCache[string, struct{}](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)
}

func (s *CacheSuite) TestCache_TwoValuesWithTheSameFrequency_ReturnTwoValues() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := InitParam{Capacity: 2}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	key1 := "key1"
	value1 := StructForCache{
		Field1: "field1",
		Field2: 100500,
	}

	key2 := "key2"
	value2 := StructForCache{
		Field1: "field2",
		Field2: 100501,
	}

	err = cache.Set(key1, value1)
	require.NoError(s.T(), err)

	err = cache.Set(key2, value2)
	require.NoError(s.T(), err)

	v1, ok := cache.Get(key1)
	require.True(s.T(), ok)
	require.Equal(s.T(), value1, v1)

	v2, ok := cache.Get(key2)
	require.True(s.T(), ok)
	require.Equal(s.T(), value2, v2)
}

func (s *CacheSuite) TestCache_ThreeValuesWithDifferentFrequency_ReturnTwoValuesAndOneWasEvicted() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := InitParam{Capacity: 2}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	key1 := "key1"
	value1 := StructForCache{
		Field1: "field1",
		Field2: 100501,
	}

	key2 := "key2"
	value2 := StructForCache{
		Field1: "field2",
		Field2: 100502,
	}

	err = cache.Set(key1, value1)
	require.NoError(s.T(), err)

	err = cache.Set(key2, value2)
	require.NoError(s.T(), err)

	v1, ok := cache.Get(key1)
	require.True(s.T(), ok)
	require.Equal(s.T(), value1, v1)

	key3 := "key3"
	value3 := StructForCache{
		Field1: "field3",
		Field2: 100503,
	}

	err = cache.Set(key3, value3)
	require.NoError(s.T(), err)

	v2, ok := cache.Get(key2)
	require.False(s.T(), ok)
	require.Equal(s.T(), StructForCache{}, v2)

	v3, ok := cache.Get(key3)
	require.True(s.T(), ok)
	require.Equal(s.T(), value3, v3)
}

func (s *CacheSuite) TestCache_SetValueByTheSameKey_ValueWasUpdated() {
	type StructForCache struct {
		Field1 string
		Field2 int
	}

	params := InitParam{Capacity: 2}
	cache, err := NewCache[string, StructForCache](params)
	require.NotNil(s.T(), cache)
	require.NoError(s.T(), err)

	key1 := "key1"
	value1 := StructForCache{
		Field1: "field1",
		Field2: 100500,
	}

	value2 := StructForCache{
		Field1: "field2",
		Field2: 100501,
	}

	err = cache.Set(key1, value1)
	require.NoError(s.T(), err)

	v1, ok := cache.Get(key1)
	require.True(s.T(), ok)
	require.Equal(s.T(), value1, v1)

	err = cache.Set(key1, value2)
	require.NoError(s.T(), err)

	v2, ok := cache.Get(key1)
	require.True(s.T(), ok)
	require.Equal(s.T(), value2, v2)
}
