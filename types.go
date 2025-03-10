package main

import (
	"cmp"
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

/*
constraints.Integer: Includes all integer types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64).
constraints.Float: Includes all floating-point types (float32, float64).
*/
type Numeric interface {
	constraints.Integer | constraints.Float
}

// Linked extends the behavior of linked list types
type Linked interface {
	GetHead() any // TODO
	GetTail() any // TODO
}

// ListNode is a generic type
type ListNode[T cmp.Ordered] struct {
	Val  T
	Next *ListNode[T]
}

// Tree used for binary tree algorithms
type Tree[T comparable] struct {
	Val   T
	Left  *Tree[T]
	Right *Tree[T]
}

// Index (generic function) returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// ErrNegativeSqrt handles negative square root errors
type ErrNegativeSqrt float64

func (n ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt of a negative number is not supported: %f", float64(n))
}

// Sqrt approximates the square root of a positive number
func Sqrt(x float64) (float64, string) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x).Error()
	}

	z := 1.0
	previousZ := 0.0

	for math.Abs(z-previousZ) > 0.000001 {
		previousZ = z
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("%f\n", z)
	}

	return z, ""
}
