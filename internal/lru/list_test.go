package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AgeListSuite struct {
	suite.Suite
}

func TestAgeListSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AgeListSuite))
}

func (s *AgeListSuite) TestNewQueue_ReturnAgeList() {
	correctCapacity := 50
	q := newAgeList[int](correctCapacity)
	require.NotNil(s.T(), q)
	assert.Equal(s.T(), correctCapacity, cap(q.data))
}

func (s *AgeListSuite) TestAdd_ValueWasAdded() {
	q := newAgeList[int](50)
	require.NotNil(s.T(), q)

	q.Add(100500)
	assert.Equal(s.T(), []int{100500}, q.data)
}

func (s *AgeListSuite) TestRemoveTheOldestValue_ListContainsValues_TheOldestValueWasRemoved() {
	q := newAgeList[int](50)
	require.NotNil(s.T(), q)

	q.Add(100500)
	q.Add(100501)
	q.Add(100502)

	keyToRemove := q.GetOldest()
	q.Remove(keyToRemove)
	assert.Equal(s.T(), []int{100501, 100502}, q.data)
}

func (s *AgeListSuite) TestRemoveTheOldestValue_ListIsEmpty_ReturnDefaultValue() {
	q := newAgeList[int](50)
	require.NotNil(s.T(), q)

	keyToRemove := q.GetOldest()
	require.Equal(s.T(), 0, keyToRemove)

	assert.NotPanics(s.T(), func() {
		q.Remove(keyToRemove)
	})
}

func (s *AgeListSuite) TestMakeValueYoungest_ValueWasMadeYoungest() {
	q := newAgeList[int](50)
	require.NotNil(s.T(), q)

	q.Add(100500)
	q.Add(100501)
	q.Add(100502)
	q.MakeYoungest(100501)
	assert.Equal(s.T(), []int{100500, 100502, 100501}, q.data)
}

func (s *AgeListSuite) TestRemove_ValueWasRemoved() {
	q := newAgeList[int](50)
	require.NotNil(s.T(), q)

	q.Add(100500)
	q.Add(100501)
	q.Add(100502)
	q.Remove(100501)
	assert.Equal(s.T(), []int{100500, 100502}, q.data)
}
