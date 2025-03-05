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
