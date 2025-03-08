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

// ListNode is a generic type
type ListNode[T cmp.Ordered] struct {
	Val  T
	Next *ListNode[T]
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
