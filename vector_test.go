package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	v := Vector2D{X: 3, Y: 5}
	assert.Equal(t, 8, Add(v))
}

func TestReduceVectors(t *testing.T) {
	vectors := []Vector2D{{1, 2}, {3, 4}, {5, 6}}
	expected := []int{3, 7, 11}
	result := ReduceVectors(vectors, Add)

	assert.Equal(t, expected, result)
}
