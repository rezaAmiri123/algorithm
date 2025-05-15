package main

import (
	"fmt"
	"sort"
)

// Great choice! The Greedy algorithm strategy is often used for optimization problems.
// Let's go over a classic Greedy example that's commonly asked in interviews:

// 🧩 Problem: Minimum Number of Coins
// Given:
// An amount of money and a list of coin denominations (e.g., [1, 5, 10, 25]),
// return the minimum number of coins required to make up that amount.

// Assumption:
// You can use infinite coins of each denomination.

// ✅ Greedy Approach:
// Always take the largest possible denomination first that doesn’t exceed the remaining amount.

// Example Input:
// coins := []int{1, 5, 10, 25}
// amount := 63
// Expected Output:
// Minimum coins: 6
// (25 + 25 + 10 + 1 + 1 + 1)

// 💡 Greedy Explanation:
// Sort the coins in descending order.
// Start from the largest coin.
// While that coin fits in the amount, use it and reduce the amount.
// Repeat with smaller coins.

func minCoins(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins))) // Descending order

	count := 0
	for _, coin := range coins {
		if amount == 0 {
			break
		}
		if coin <= amount {
			num := amount / coin
			count += num
			amount -= coin * num
		}
	}
	if amount != 0 {
		return -1 // not possible to make the amount
	}
	return count
}

func main() {
	coins := []int{1, 5, 10, 25}
	amount := 63
	fmt.Printf("Minimum coins: %d\n", minCoins(coins, amount))
}

// 📝 Note:
// This greedy approach works correctly only for coin systems like US denominations 
// where it always leads to the optimal solution. For arbitrary coin sets, 
// dynamic programming might be needed.

// Would you like a version of this problem that needs DP because greedy fails?
