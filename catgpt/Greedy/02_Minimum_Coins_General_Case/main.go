package main

import (
	"fmt"
	"math"
)

// Great! Let’s explore a case where the Greedy algorithm fails
// and Dynamic Programming (DP) is required.

// 🧩 Problem: Minimum Coins (General Case)
// Given:
// An amount and a list of coin denominations.
// Return the minimum number of coins needed to make that amount.

// 🔥 Example Where Greedy Fails:
// coins := []int{1, 3, 4}
// amount := 6

// ❌ Greedy Approach:
// Start with the largest coin ≤ 6 → 4
// 6 - 4 = 2 → take 1 → 1 → total coins = 3
// But the optimal is: 3 + 3 = 6 → only 2 coins

func minCoinsDP(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 3, 4}
	amount := 6
	fmt.Printf("Minimum coins (DP): %d\n", minCoinsDP(coins, amount))
}

// ✅ Why DP Works:
// It explores all combinations using smaller subproblems. 
// Greedy only makes local choices without looking ahead, which can miss the optimal solution.

// Would you like a visual explanation or diagram for how the DP array fills up in this example?
