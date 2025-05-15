package main

import "fmt"

// 🧩 Problem: Generate All Subsets (Power Set)
// Given a set of distinct integers nums, return all possible subsets (the power set).

// Input:
// nums := []int{1, 2, 3}
// Output:
// [[], [1], [2], [3], [1 2], [1 3], [2 3], [1 2 3]]

// 🧠 Explanation
// This is a typical backtracking problem, where at each step,
// you decide whether to include or exclude an element.
// Start with an empty set.
// At each level in recursion, you either:
// Include the current element in the subset.
// Exclude it and move to the next.
// This creates a decision tree of all combinations.

func subsets(nums []int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// Make a copy of current and append to result
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			//choose
			current = append(current, nums[i])
			// explore
			backtrack(i + 1)
			// Unchoose
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println("All subsets:")
	for _, subset := range result {
		fmt.Println(subset)
	}
}

// 🔍 How It Works
// backtrack(start int) explores all subsets starting from index start.
// current is the current subset being built.
// Add copy of current to result each time to avoid reference issues.
// Backtracking is done by removing the last added element (Unchoose) after the recursive call.
// Would you like a harder version next, like solving the N-Queens problem in Go?
