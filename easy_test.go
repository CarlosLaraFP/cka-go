package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
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

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{""}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"", ""}))
	assert.Equal(t, "cat", LongestCommonPrefix([]string{"cat"}))
	assert.Equal(t, "a", LongestCommonPrefix([]string{"ab", "a"}))
}

// Helper function to create a linked list from a slice
func createList[T constraints.Ordered](values []T) *ListNode[T] {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode[T]{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode[T]{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice
func listToSlice[T constraints.Ordered](head *ListNode[T]) []T {
	var result []T
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{1}, []int{1}, []int{1, 1}},
	}

	for _, tt := range tests {
		list1 := createList(tt.list1)
		list2 := createList(tt.list2)
		merged := MergeTwoLists(list1, list2)
		got := listToSlice(merged)
		if !equalSlices(got, tt.want) {
			t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.want)
		}
	}
}

// Helper function to compare two slices
func equalSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
