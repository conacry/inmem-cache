package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type QueueSuite struct {
	suite.Suite
}

func TestQueueSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(QueueSuite))
}

func (s *QueueSuite) TestNewQueue_IllegalCapacity_ReturnQueue() {
	illegalCapacity := 0
	expectedErr := "capacity should be greater than 0"

	q, err := NewAgeList[int](illegalCapacity)
	assert.Nil(s.T(), q)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr, err.Error())
}

func (s *QueueSuite) TestNewQueue_CorrectCapacity_ReturnQueue() {
	correctCapacity := 50
	q, err := NewAgeList[int](correctCapacity)
	require.NotNil(s.T(), q)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), correctCapacity, cap(q.data))
}

func (s *QueueSuite) TestAdd_ValueWasAdded() {
	q, err := NewAgeList[int](50)
	require.NotNil(s.T(), q)
	require.NoError(s.T(), err)

	q.Add(100500)
	assert.Equal(s.T(), []int{100500}, q.data)
}

func (s *QueueSuite) TestRemoveTheOldestValue_TheOldestValueWasRemoved() {
	q, err := NewAgeList[int](50)
	require.NotNil(s.T(), q)
	require.NoError(s.T(), err)

	q.Add(100500)
	q.Add(100501)
	q.Add(100502)

	keyToRemove := q.GetOldest()
	q.Remove(keyToRemove)
	assert.Equal(s.T(), []int{100501, 100502}, q.data)
}

func (s *QueueSuite) TestMakeValueYoungest_ValueWasMadeYoungest() {
	q, err := NewAgeList[int](50)
	require.NotNil(s.T(), q)
	require.NoError(s.T(), err)

	q.Add(100500)
	q.Add(100501)
	q.Add(100502)
	q.MakeYoungest(100501)
	assert.Equal(s.T(), []int{100500, 100502, 100501}, q.data)
}

func (s *QueueSuite) TestRemove_ValueWasRemoved() {
	q, err := NewAgeList[int](50)
	require.NotNil(s.T(), q)
	require.NoError(s.T(), err)

	q.Add(100500)
	q.Add(100501)
	q.Add(100502)
	q.Remove(100501)
	assert.Equal(s.T(), []int{100500, 100502}, q.data)
}
