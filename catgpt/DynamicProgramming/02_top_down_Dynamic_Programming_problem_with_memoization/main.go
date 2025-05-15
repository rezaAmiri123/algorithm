package main

import "fmt"

// ✅ Example Problem: Unique Paths
// Problem Statement:
// A robot is located at the top-left corner of a m x n grid.
// The robot can only move either down or right at any point in time.
// Return the number of possible unique paths
// that the robot can take to reach the bottom-right corner.

// 🧠 Approach: Top-Down (Memoization)
// This is a classic recursive DP problem.
// At each cell (i, j), the number of unique paths to the bottom-right is the sum of:
// The paths from the cell below (i+1, j)
// The paths from the cell to the right (i, j+1)
// We use a memoization table to avoid recomputing subproblems.

// 🔢 Recurrence Relation:
// uniquePaths(i, j) = uniquePaths(i+1, j) + uniquePaths(i, j+1)
// Base Case: if i == m-1 or j == n-1, there's only 1 path (go straight).
// Use a 2D slice memo[i][j] to cache results.

func uniquePaths(m, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // initialize with -1 to mark unvisited
		}
	}
	return dfs(0, 0, m, n, memo)
}

func dfs(i, j, m, n int, memo [][]int) int {
	if i >= m || j >= n {
		return 0
	}
	if i == m-1 && j == n-1 {
		return 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	memo[i][j] = dfs(i+1, j, m, n, memo) + dfs(i, j+1, m, n, memo)
	return memo[i][j]
}

func main() {
	m, n := 3, 7
	fmt.Printf("Unique Paths in a %dx%d grid: %d\n", m, n, uniquePaths(m, n)) // Output: 28
}

// 📝 Explanation:
// From the top-left (0,0) to bottom-right (2,6), 
// the robot can make 28 unique paths by only moving right or down.
// We avoid recomputation by storing results in memo[i][j].
// Would you like to try a DP with 1D space optimization, 
// or a real interview problem like "Word Break" or "Knapsack" next?
