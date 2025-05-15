package main

import "fmt"

// ✅ Example Problem: Climbing Stairs
// Problem Statement:
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can climb either 1 or 2 steps.
// Return the number of distinct ways to climb to the top.

// 🧠 Approach: Bottom-Up with 1D DP
// This is a classic Fibonacci-style DP problem.
// Instead of using a full array, we only need the last two values,
// so we can optimize to O(1) space using just two variables.

// 🔢 Recurrence:
// dp[i] = dp[i-1] + dp[i-2]
// We only need to store prev1 (dp[i-1]) and prev2 (dp[i-2]).

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	perv1, perv2 := 2, 1 // dp[2] = 2, dp[1] = 1
	for i := 3; i <= n; i++ {
		curr := perv1 + perv2
		perv2 = perv1
		perv1 = curr
	}

	return perv1
}

func main() {
	n := 5
	fmt.Printf("Number of ways to climb %d stairs: %d\n", n, climbStairs(n)) // Output: 8
}

// 📚 Explanation:
// n = 5 stairs
// Ways: [1,1,1,1,1], [1,1,1,2], [1,2,2], [2,1,2], [2,2,1], etc.
// Total = 8 unique ways

// This kind of optimization is commonly used in:
// Fibonacci problems
// House Robber (when you only need i-1 and i-2)
// Optimized Knapsack (when only row i-1 matters)
// Would you like to try a 1D-optimized version of the "House Robber" problem next?
