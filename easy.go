package main

import (
	"cmp"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. In Go, we use slices.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
Time complexity: O(N) | Space complexity: Amortized O(N)
*/
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for b, n := range nums {
		remainder := target - n

		if a, exists := numMap[remainder]; exists {
			return []int{a, b}
		}

		numMap[n] = b
	}

	return nil
}

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Time complexity: O(N) | Space complexity: O(N) worst case (all open brackets keeps growing the stack)
*/
func IsValid(s string) bool {
	stack := []rune{}

	// Mapping closing brackets to their corresponding opening brackets
	matches := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if open, exists := matches[char]; exists {
			// Check if the stack is empty or the top is not the correct opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Pop the last element
			stack = stack[:len(stack)-1]
		} else {
			// Otherwise, it's an opening bracket since all map keys are closing brackets; push it to the stack
			stack = append(stack, char)
		}
	}

	// Stack should be empty if all brackets matched correctly
	return len(stack) == 0
}

/*
Given a roman numeral, convert it to an integer.
For example, 2 is written as II in Roman numeral, just two ones added together.
12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written as IV.
Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX.
There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9 (2 extra in the naive sum in both cases).
X can be placed before L (50) and C (100) to make 40 and 90 (20 extra in the naive sum in both cases).
C can be placed before D (500) and M (1000) to make 400 and 900 (200 extra in the naive sum in both cases).

Time complexity: O(N)
Space complexity: O(1)
*/
func RomanToInt(s string) int {
	dict := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0

	for i, currentChar := range s {
		sum += dict[currentChar]

		if i > 0 {
			if previousChar := rune(s[i-1]); dict[previousChar] < dict[currentChar] {
				sum -= dict[previousChar] * 2
			}
		}
	}

	return sum
}

/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Time complexity: O(N * M) worst case
Space complexity: O(N)
*/
func LongestCommonPrefix(strs []string) string {
	wordIndex := 0
	charIndex := 0
	var commonPrefix []byte
	var currentChar byte = 0

	for {
		word := strs[wordIndex]
		if charIndex >= len(word) {
			return string(commonPrefix)
		}
		char := word[charIndex]

		if currentChar == 0 {
			currentChar = char
		} else if currentChar != char {
			return string(commonPrefix)
		}

		wordIndex++

		// if we already went through all the words, proceed to the next character
		if wordIndex == len(strs) {
			commonPrefix = append(commonPrefix, currentChar)
			wordIndex = 0
			charIndex++
			currentChar = 0
		}
	}
}

/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/
func MergeTwoLists[T cmp.Ordered](list1 *ListNode[T], list2 *ListNode[T]) *ListNode[T] {
	// Create a dummy node to simplify the merging process
	dummy := &ListNode[T]{}
	current := dummy

	// Traverse both lists and merge them
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Attach the remaining nodes from list1 or list2
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the head of the merged list
	return dummy.Next
}

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially.
The remaining elements of nums are not important as well as the size of nums. Return k.
*/
func RemoveDuplicates(nums []int) int {
	// 1) always compare the current value and the previous value
	// 2) if the numbers are the same, stop incrementing the replacement index
	// 3) if the numbers are different, replace the index value with the current value
	// Example: 1, 2, 2, 3
	index := 1 // minimum starting replacement index

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i] // index and i are the same under normal sequence
			index++
		}
		// if previous == current, do not increment index
	}

	return index
}

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
Input: haystack = "sadbutsad", needle = "sad" | Output: 0
Input: haystack = "leetcode", needle = "leeto" | Output: -1

	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
*/
func FirstOcurrence(haystack string, needle string) int {
	// Refer to simple solution above
	haystackIndex := 0
	needleIndex := 0
	firstIndex := -1

	for haystackIndex < len(haystack) && needleIndex < len(needle) {
		if haystack[haystackIndex] != needle[needleIndex] {
			haystackIndex = haystackIndex - needleIndex // reset index for a new attempt starting 1 char ahead
			needleIndex = 0                             // restart search
			firstIndex = -1                             // not found
		} else {
			// set a new index only if we can still fit in the needle
			if len(haystack)-haystackIndex >= len(needle) {
				firstIndex = haystackIndex - needleIndex
			}
			needleIndex++
		}
		haystackIndex++ // haystack index always increments
	}

	return firstIndex
}

/*
You are given a large integer represented as an integer array digits, where each digits[i]
is the ith digit of the integer. The digits are ordered from most significant to least
significant in left-to-right order. The large integer does not contain any leading 0's.
Increment the large integer by one and return the resulting array of digits.
*/
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if i == 0 && digits[i] == 9 {
			digits[i] = 1
			digits = append(digits, 0)
		} else if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}

	return digits
}

/*
The square root of a number is the number which when multiplied by itself gives the first number.
Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.
*/
func MySqrt(x int) int {
	// no floats at all
	start, end := 0, x+1 // it doesn't hurt to extend the end slightly in exchange for taking care of the zero case
	// when start equals end, no more progress can be made
	for start < end {
		mid := start + (end-start)/2

		if mid*mid > x {
			end = mid
		} else {
			start = mid + 1 // otherwise keeps bouncing back and forth
		}
	}

	return start - 1 // always an integer higher
}

/*
TreeDepth finds the depth of a binary tree.
*/
func TreeDepth[T comparable](root *Tree[T]) int {
	if root == nil {
		return 0
	}

	leftDepth := TreeDepth(root.Left)
	rightDepth := TreeDepth(root.Right)

	// + 1 accounts for the current node
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// go test ./... -v
