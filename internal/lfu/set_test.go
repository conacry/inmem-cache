package lfucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FrequencySetSuite struct {
	suite.Suite
}

func TestFrequencySetSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(FrequencySetSuite))
}

func (s *FrequencySetSuite) TestNewFrequencySet_ReturnFrequencySet() {
	set := newFrequencySet[string]()
	require.NotNil(s.T(), set)
	assert.NotNil(s.T(), set.data)
	assert.Equal(s.T(), 0, set.minCount)
}

func (s *FrequencySetSuite) TestAdd_NewValueWasAdded() {
	set := newFrequencySet[string]()
	require.NotNil(s.T(), set)

	value := "key1"
	set.Add(value)

	bucket, ok := set.data[0]
	assert.True(s.T(), ok)
	assert.NotNil(s.T(), bucket)

	createdAt, ok := set.data[0][value]
	assert.True(s.T(), ok)
	assert.NotEmpty(s.T(), createdAt)
}

func (s *FrequencySetSuite) TestTouch_FrequencyCountIncreased() {
	set := newFrequencySet[string]()
	require.NotNil(s.T(), set)

	value := "key1"
	set.Add(value)
	bucket, ok := set.data[0]
	require.True(s.T(), ok)
	assert.NotNil(s.T(), bucket)

	_, ok = set.data[0][value]
	require.True(s.T(), ok)

	set.Touch(value, 0)

	_, ok = set.data[0][value]
	assert.False(s.T(), ok)

	_, ok = set.data[1][value]
	assert.True(s.T(), ok)
}

func (s *FrequencySetSuite) TestGetLeastFrequent_ReturnLessUsedValue() {
	set := newFrequencySet[string]()
	require.NotNil(s.T(), set)

	valueOne := "key1"
	countOne := 0
	set.Add(valueOne)

	_, ok := set.data[countOne][valueOne]
	require.True(s.T(), ok)

	set.Touch(valueOne, countOne)
	countOne++
	set.Touch(valueOne, countOne)
	countOne++
	set.Touch(valueOne, countOne)

	valueTwo := "key2"
	countTwo := 0
	set.Add(valueTwo)

	_, ok = set.data[countTwo][valueTwo]
	require.True(s.T(), ok)

	set.Touch(valueTwo, 0)
	set.Touch(valueTwo, 1)

	leastFrequent := set.GetLeastFrequent()
	assert.Equal(s.T(), valueTwo, leastFrequent)
}

func (s *FrequencySetSuite) TestRemove_ValueWasRemoved() {
	set := newFrequencySet[string]()
	require.NotNil(s.T(), set)

	valueOne := "key1"
	countOne := 0
	set.Add(valueOne)

	_, ok := set.data[countOne][valueOne]
	require.True(s.T(), ok)

	set.Touch(valueOne, countOne)
	countOne++
	set.Touch(valueOne, countOne)
	countOne++
	set.Touch(valueOne, countOne)

	valueTwo := "key2"
	countTwo := 0
	set.Add(valueTwo)

	_, ok = set.data[countTwo][valueTwo]
	require.True(s.T(), ok)

	set.Touch(valueTwo, 0)
	set.Touch(valueTwo, 1)

	leastFrequent := set.GetLeastFrequent()
	assert.Equal(s.T(), valueTwo, leastFrequent)

	set.Remove(valueTwo, countTwo)
	_, ok = set.data[countTwo][valueTwo]
	assert.False(s.T(), ok)
}
