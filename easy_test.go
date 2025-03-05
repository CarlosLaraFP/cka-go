package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, TwoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, TwoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, TwoSum([]int{3, 3}, 6))
}

func TestIsValid(t *testing.T) {
	assert.Equal(t, false, IsValid("}{"))
	assert.Equal(t, true, IsValid("{}"))
	assert.Equal(t, true, IsValid("{}()[]"))
	assert.Equal(t, true, IsValid("{([])}"))
	assert.Equal(t, false, IsValid("{([]})"))
	assert.Equal(t, false, IsValid("((("))
}

func TestRomanToInt(t *testing.T) {
	assert.Equal(t, 3, RomanToInt("III"))
	assert.Equal(t, 4, RomanToInt("IV"))
	assert.Equal(t, 5, RomanToInt("V"))
	assert.Equal(t, 9, RomanToInt("IX"))
	assert.Equal(t, 40, RomanToInt("XL"))
	assert.Equal(t, 90, RomanToInt("XC"))
	assert.Equal(t, 400, RomanToInt("CD"))
	assert.Equal(t, 900, RomanToInt("CM"))
}
