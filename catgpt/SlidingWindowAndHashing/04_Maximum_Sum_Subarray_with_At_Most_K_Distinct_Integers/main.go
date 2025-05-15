package main

import "fmt"

// 🔍 Problem: Maximum Sum Subarray with At Most K Distinct Integers
// Given an array of integers nums and an integer k,
// find the maximum sum of any subarray that contains at most k distinct integers.

// Example:
// Input: nums = [1, 2, 1, 2, 3], k = 2
// Output: 6
// Explanation: The subarray [1, 2, 1, 2] has 2 distinct integers and sum = 6.

// 💡 Explanation
// We combine sliding window + hash map + running sum.

// Key additions:
// Track current window sum.
// Update maxSum only when distinct count ≤ k.

func maxSumSubarrayWithKDistinct(nums []int, k int) int {
	if k == 0 || len(nums) == 0 {
		return 0
	}

	countMap := make(map[int]int)
	start := 0
	currentSum := 0
	maxSum := 0

	for end := 0; end < len(nums); end++ {
		countMap[nums[end]]++
		currentSum += nums[end]

		for len(countMap) > k {
			countMap[nums[start]]--
			currentSum -= nums[start]
			if countMap[nums[start]] == 0 {
				delete(countMap, nums[start])
			}
			start++
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{1, 2, 1, 2, 3}
	k := 2
	fmt.Println("Maximum sum of subarray with at most", k, "distinct integers:", maxSumSubarrayWithKDistinct(nums, k))
}

// 🧠 Interview Highlights:
// Introduces sum tracking along with distinct count.
// Efficient: O(n) time and O(k) space.
// Great for showing fluency with multiple patterns (hashing + window + math).

// Would you like a challenge version where it requires exactly k distinct integers 
// and maximum sum?
