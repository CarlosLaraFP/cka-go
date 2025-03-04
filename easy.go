package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
Time complexity: O(N) | Space complexity: O(N)
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
