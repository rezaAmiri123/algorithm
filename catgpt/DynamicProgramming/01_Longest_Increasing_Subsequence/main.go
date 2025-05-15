package main

import "fmt"

// ✅ Example Problem: Longest Increasing Subsequence (LIS)
// Problem Statement:
// Given an integer array nums, return the length of the longest strictly increasing subsequence.

// 🧠 Approach:
// We'll use bottom-up dynamic programming.
// Let dp[i] represent the length of the longest increasing subsequence ending at index i.
// For every element nums[i], we check all previous elements nums[j] (where j < i),
// and if nums[j] < nums[i], we can extend the subsequence ending at j to i.
// So the recurrence is:
// dp[i] = max(dp[i], dp[j]+1) if nums[j] < nums[i]
// ⏱ Time Complexity:
// O(n^2) where n is the length of nums.

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // every number is at least a subsequence of length 1
	}

	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; i < j; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Length of LIS:", lengthOfLIS(nums)) // Output: 4 (e.g., 2, 3, 7, 101)
}

// 📚 Explanation:
// For the input [10, 9, 2, 5, 3, 7, 101, 18], the longest increasing subsequence is [2, 3, 7, 101], which has length 4.

// Would you like a harder DP problem next, or one that uses memoization instead of tabulation?