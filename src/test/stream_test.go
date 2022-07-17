package test

import (
	"github.com/stretchr/testify/assert"
	"golambdas/src/golambdas"
	"testing"
)

func TestStreamFilter(t *testing.T) {
	assert := assert.New(t)
	var golambdas = golambdas.NewStream([]int{1, 2, 3, 4, 5, -24356756585342})

	var result = golambdas.Filter(func(element int) bool {
		return element > 2
	})

	assert.Lenf(result, 3, "Expected 3 elements, got %d", len(result))
}

func TestStreamAllMatch(t *testing.T) {
	assert := assert.New(t)
	var golambdas = golambdas.NewStream([]int{1, 2, 3, 4, 5})

	var result = golambdas.AllMatch(func(element int) bool {
		return element > 2
	})

	assert.Falsef(result, "Expected false, got %t", result)
}

func TestStreamForEach(t *testing.T) {
	assert := assert.New(t)
	var golambdas = golambdas.NewStream([]int{1, 2, 3, 4, 5})

	counter := 0
	golambdas.ForEach(func(element int) {
		counter += element
	})

	assert.Equalf(counter, 15, "Expected 15, got %d", counter)
}

func TestStreamReduce(t *testing.T) {
	assert := assert.New(t)
	var golambdas = golambdas.NewStream([]int{1, 2, 3, 4, 5})

	var result = golambdas.Reduce(10, func(element int, accumulator int) int {
		return accumulator + element
	})

	assert.Equalf(result, 25, "Expected 15, got %d", result)
}
