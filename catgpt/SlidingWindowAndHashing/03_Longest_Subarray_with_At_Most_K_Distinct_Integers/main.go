package main

import "fmt"

// 🔍 Problem: Longest Subarray with At Most K Distinct Integers
// Given an integer array nums and an integer k,
// return the length of the longest subarray that contains at most k distinct integers.

// Example:
// Input: nums = [1, 2, 1, 2, 3], k = 2
// Output: 4
// Explanation: The subarray [1, 2, 1, 2] has 2 distinct integers.

// 💡 Explanation
// It’s the same idea as before but applied to an array of integers.

// Steps:
// Use a map[int]int to track the count of each integer.
// Maintain a sliding window with start and end.
// Expand end, updating the map.
// When the map size exceeds k, shrink from the start.
// Update max length each time.

func longestSubarrayWithKDistinct(nums []int, k int) int {
	if k == 0 || len(nums) == 0 {
		return 0
	}

	countMap := make(map[int]int)
	start := 0
	maxLen := 0

	for end := 0; end < len(nums); end++ {
		countMap[nums[end]]++

		for len(countMap) > k {
			countMap[nums[start]]--
			if countMap[nums[start]] == 0 {
				delete(countMap, nums[start])
			}
			start++
		}
		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}
	return maxLen
}

func main() {
	nums := []int{1, 2, 1, 2, 3}
	k := 2
	fmt.Println("Length of longest subarray with at most", k, "distinct integers:", longestSubarrayWithKDistinct(nums, k))
}

// 🧠 Interview Value:
// Maps to real-world scenarios (e.g. session tracking, sensor data, etc).
// Demonstrates understanding of:
// Sliding window logic.
// Frequency counting in hash maps.
// Edge case handling.

// Would you like to try the at most K distinct + maximum sum version next, 
// or a version where exactly K distinct is required?
