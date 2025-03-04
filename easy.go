package main

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
