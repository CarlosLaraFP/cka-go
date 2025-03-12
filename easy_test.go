package main

import (
	"testing"

	"cmp"

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

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{""}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"", ""}))
	assert.Equal(t, "cat", LongestCommonPrefix([]string{"cat"}))
	assert.Equal(t, "a", LongestCommonPrefix([]string{"ab", "a"}))
}

// Helper function to create a linked list from a slice
func createList[T cmp.Ordered](values []T) *ListNode[T] {
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
func listToSlice[T cmp.Ordered](head *ListNode[T]) []T {
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

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 1, RemoveDuplicates([]int{-1}))
	assert.Equal(t, 3, RemoveDuplicates([]int{-1, 0, 1}))
	assert.Equal(t, 3, RemoveDuplicates([]int{1, 2, 2, 3}))
	assert.Equal(t, 5, RemoveDuplicates([]int{1, 2, 2, 3, 4, 5}))
	assert.Equal(t, 5, RemoveDuplicates([]int{1, 2, 2, 3, 4, 4, 5}))
	assert.Equal(t, 5, RemoveDuplicates([]int{1, 2, 2, 3, 4, 4, 5, 5, 5}))
}

func TestFirstOcurrence(t *testing.T) {
	assert.Equal(t, -1, FirstOcurrence("haystack", "needle"))
	assert.Equal(t, 0, FirstOcurrence("leetcode", "leet"))
	assert.Equal(t, 0, FirstOcurrence("leetcodeleetcode", "leet"))
	assert.Equal(t, -1, FirstOcurrence("leet", "leetcode"))
	assert.Equal(t, 4, FirstOcurrence("mississippi", "issip"))
	assert.Equal(t, -1, FirstOcurrence("mississippi", "issipi"))
}

// 9999 -> 10000
func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5}, PlusOne([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 4, 0}, PlusOne([]int{1, 2, 3, 9}))
	assert.Equal(t, []int{1, 0, 0, 0, 0}, PlusOne([]int{9, 9, 9, 9}))
	assert.Equal(t, []int{1, 1, 0, 0}, PlusOne([]int{1, 0, 9, 9}))
	assert.Equal(t, []int{1, 0}, PlusOne([]int{9}))
}

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 2, MySqrt(4))
	assert.Equal(t, 2, MySqrt(8))
	assert.Equal(t, 5, MySqrt(27))
	assert.Equal(t, 9, MySqrt(81))
	assert.Equal(t, 0, MySqrt(0))
	assert.Equal(t, 15740, MySqrt(247776352))
	assert.Equal(t, 46339, MySqrt(2147395599))
}

func TestTreeDepth(t *testing.T) {
	a := &Tree[int]{Val: 5, Left: nil, Right: nil}
	assert.Equal(t, 1, TreeDepth(a))
	b := &Tree[int]{Val: 5, Left: &Tree[int]{Val: 5, Left: nil, Right: nil}, Right: nil}
	assert.Equal(t, 2, TreeDepth(b))
	c := &Tree[int]{Val: 5, Left: a, Right: b}
	assert.Equal(t, 3, TreeDepth(c))
	d := &Tree[int]{
		Val: 1,
		Left: &Tree[int]{
			Val:   2,
			Left:  &Tree[int]{Val: 4, Left: a},
			Right: &Tree[int]{Val: 5},
		},
		Right: &Tree[int]{Val: 3, Right: b},
	}
	assert.Equal(t, 4, TreeDepth(d))
}

func TestInorderTraversal(t *testing.T) {
	a := &Tree[int]{Val: 1, Left: nil, Right: nil}
	assert.Equal(t, []int{1}, InorderTraversal(a))
	assert.Equal(t, []int{1}, InorderTraversalIterative(a))
	b := &Tree[int]{Val: 5, Left: &Tree[int]{Val: 4}}
	assert.Equal(t, []int{4, 5}, InorderTraversal(b))
	assert.Equal(t, []int{4, 5}, InorderTraversalIterative(b))
	c := &Tree[int]{Val: 2, Left: a, Right: b}
	assert.Equal(t, []int{1, 2, 4, 5}, InorderTraversal(c))
	assert.Equal(t, []int{1, 2, 4, 5}, InorderTraversalIterative(c))
	d := &Tree[int]{
		Val: 1,
		Left: &Tree[int]{
			Val:   2,
			Left:  &Tree[int]{Val: 4, Left: a},
			Right: &Tree[int]{Val: 5},
		},
		Right: &Tree[int]{Val: 3, Right: b},
	}
	assert.Equal(t, []int{1, 4, 2, 5, 1, 3, 4, 5}, InorderTraversal(d))
	assert.Equal(t, []int{1, 4, 2, 5, 1, 3, 4, 5}, InorderTraversalIterative(d))
}

func TestIsSymmetric(t *testing.T) {
	a := &Tree[int]{Val: 1}
	assert.Equal(t, true, IsSymmetric(a))
	assert.Equal(t, true, IsSymmetricIterative(a))
	b := &Tree[int]{Val: 5, Left: &Tree[int]{Val: 4}}
	assert.Equal(t, false, IsSymmetric(b))
	assert.Equal(t, false, IsSymmetricIterative(b))
	c := &Tree[int]{Val: 2, Left: a, Right: &Tree[int]{Val: 1}}
	assert.Equal(t, true, IsSymmetric(c))
	assert.Equal(t, true, IsSymmetricIterative(c))
	d := &Tree[int]{
		Val: 1,
		Left: &Tree[int]{
			Val:   2,
			Left:  &Tree[int]{Val: 4, Left: a},
			Right: &Tree[int]{Val: 5},
		},
		Right: &Tree[int]{Val: 3, Right: b},
	}
	assert.Equal(t, false, IsSymmetric(d))
	assert.Equal(t, false, IsSymmetricIterative(d))
	e := &Tree[int]{
		Val: 1,
		Left: &Tree[int]{
			Val:   2,
			Left:  &Tree[int]{Val: 4, Left: &Tree[int]{Val: 1}},
			Right: &Tree[int]{Val: 5},
		},
		Right: &Tree[int]{
			Val:   2,
			Left:  &Tree[int]{Val: 5},
			Right: &Tree[int]{Val: 4, Right: &Tree[int]{Val: 1}},
		},
	}
	assert.Equal(t, true, IsSymmetric(e))
	assert.Equal(t, true, IsSymmetricIterative(e))
}

func TestSameTrees(t *testing.T) {
	a := &Tree[int]{Val: 1}
	assert.Equal(t, true, SameTrees(a, &Tree[int]{Val: 1}))
	b := &Tree[int]{Val: 5, Left: &Tree[int]{Val: 4}}
	assert.Equal(t, false, SameTrees(a, b))
	c := &Tree[int]{Val: 2, Left: a, Right: &Tree[int]{Val: 1}}
	assert.Equal(t, false, SameTrees(c, b))
	e := &Tree[int]{
		Val: 1,
		Left: &Tree[int]{
			Val:   2,
			Left:  &Tree[int]{Val: 4, Left: &Tree[int]{Val: 1}},
			Right: &Tree[int]{Val: 5},
		},
		Right: &Tree[int]{
			Val:   2,
			Left:  &Tree[int]{Val: 4, Left: &Tree[int]{Val: 1}},
			Right: &Tree[int]{Val: 5},
		},
	}
	assert.Equal(t, true, SameTrees(e.Left, e.Right))
}
