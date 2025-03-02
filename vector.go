package main

// Vector2D represents a tuple of two numbers
type Vector2D struct {
	X int
	Y int
}

// Add two numbers
func Add(v Vector2D) int {
	return v.X + v.Y
}

// ReduceVectors applies the provided operation on each Vector2D in the slice and returns the results.
func ReduceVectors(vectors []Vector2D, operation func(Vector2D) int) []int {
	results := make([]int, len(vectors))
	for i, v := range vectors {
		results[i] = operation(v)
	}
	return results
}
