package main

import (
	"fmt"
	"sort"
)

// ✅ Problem: Merge Intervals
// Description:
// Given an array of intervals where intervals[i] = [start_i, end_i],
// merge all overlapping intervals, and return an array of the non-overlapping intervals
// that cover all the intervals in the input.

// 🔍 Example
// Input:
// intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
// Output:
// [][]int{{1,6},{8,10},{15,18}}

// Explanation:
// [1,3] and [2,6] overlap, so merge them into [1,6].
// [8,10] and [15,18] do not overlap with anything.

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals by starting time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	current := intervals[0]

	for _, next := range intervals[1:] {
		if next[0] <= current[1] { // overlap
			current[1] = max(current[1], next[1])
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge(intervals)
	fmt.Println(result) // Output: [[1 6] [8 10] [15 18]]
}

// Would you like to try a more difficult interval problem like "Insert Interval" 
// or "Meeting Rooms"?