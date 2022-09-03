package going

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStream_NoElements(t *testing.T) {
	// act
	stream := NewStream[int]()

	// assert
	assert.NotNil(t, stream)
	for range stream {
		assert.Fail(t, "Stream must be empty")
	}
}

func TestNewStream_OneElement(t *testing.T) {
	// act
	stream := NewStream(1)

	// assert
	assert.NotNil(t, stream)
	for el := range stream {
		assert.Equal(t, 1, el)
	}
}

func TestStream_Filter_SomeExcluded(t *testing.T) {
	// arrange
	stream := NewStream(1, 2, 3, 4)
	isEven := func(x int) bool { return x%2 == 0 }

	// act
	actual := stream.Filter(isEven).ToSlice()

	// assert
	assert.Equal(t, []int{2, 4}, actual)
}

func TestStream_Filter_AllExcluded(t *testing.T) {
	// arrange
	stream := NewStream(1, 2, 3, 4)
	none := func(x int) bool { return false }

	// act
	actual := stream.Filter(none).ToSlice()

	// assert
	assert.Equal(t, []int{}, actual)
}

func TestStream_Map(t *testing.T) {
	// arrange
	stream := NewStream(1, 2, 3, 4)
	double := func(x int) int { return x * 2 }

	// act
	actual := stream.Map(double).ToSlice()

	// assert
	assert.Equal(t, []int{2, 4, 6, 8}, actual)
}

func TestStream_Map_NoElements(t *testing.T) {
	// arrange
	stream := NewStream[int]()
	double := func(x int) int { return x * 2 }

	// act
	actual := stream.Map(double).ToSlice()

	// assert
	assert.Equal(t, []int{}, actual)
}

func TestStream_Count(t *testing.T) {
	// arrange
	stream := NewStream(10, 20, 30, 40)

	// act
	actual := stream.Count()

	// assert
	assert.Equal(t, 4, actual)
}

func TestStream_Count_Empty(t *testing.T) {
	// arrange
	stream := NewStream[string]()

	// act
	actual := stream.Count()

	// assert
	assert.Equal(t, 0, actual)
}
