package main

import "fmt"

// ✅ Problem: House Robber
// Problem Statement:
// You are given an integer array nums representing the amount of money
// in each house along a street.
// You cannot rob two adjacent houses (they will alert the police).
// Return the maximum amount of money you can rob without alerting the police.

// 🧠 Approach: Bottom-Up DP with 1D Space
// For each house i, you have two options:
// Rob it: then you must skip the previous house (i-1), so add nums[i] to the result from i-2.
// Skip it: take the result from i-1.
// So, recurrence:
// dp[i] = max(dp[i-1], dp[i-2] + nums[i])

// But we can optimize space to use just two variables: prev1 and prev2.

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	prev2 := nums[0]               // dp[i-2]
	prev1 := max(nums[0], nums[1]) // dp[i-1]

	for i := 2; i < n; i++ {
		curr := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println("Max money you can rob:", rob(nums)) // Output: 12
}

// 📚 Explanation:
// Houses: [2, 7, 9, 3, 1]
// Rob 2 + 9 + 1 = 12 (skip adjacent 7 and 3)

// This is the maximum possible without robbing two adjacent houses.
// This problem often leads into advanced variations like House Robber II (circular street) 
// or Tree-based versions.
// Would you like to try House Robber II next, or something more complex like 0/1 Knapsack?
