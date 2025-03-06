package main

import "golang.org/x/exp/constraints"

/*
constraints.Integer: Includes all integer types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64).
constraints.Float: Includes all floating-point types (float32, float64).
*/
type Numeric interface {
	constraints.Integer | constraints.Float
}

// ListNode is a generic type
type ListNode[T constraints.Ordered] struct {
	Val  T
	Next *ListNode[T]
}
