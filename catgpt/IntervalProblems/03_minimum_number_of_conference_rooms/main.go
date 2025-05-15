package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// ✅ Problem: Meeting Rooms II
// Description:
// Given an array of meeting time intervals [[s1,e1],[s2,e2],...],
// return the minimum number of conference rooms required.

// 🔍 Example
// Input:
// intervals := [][]int{{0,30},{5,10},{15,20}}
// Output:
// Explanation:
// [0,30] and [5,10] overlap → need 2 rooms.
// [15,20] overlaps with [0,30], so still need 2 rooms total.

// 💡 Approach
// We can solve this efficiently using a min-heap (priority queue):
// Sort the intervals by start time.
// Use a min-heap to track the end times of ongoing meetings.

// For each interval:
// If the current meeting starts after or at the end time of the earliest meeting,
// reuse that room (pop from heap).
// Otherwise, need a new room (push to heap).
// The size of the heap at the end tells us the minimum number of rooms needed.

// A MinHeap to store meeting end times
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[0 : n-1]
	return val
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	// Add the end time of the first meeting
	heap.Push(h, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		// If the earliest room is free, reuse it
		if intervals[i][0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, intervals[i][1])
	}

	return h.Len()
}

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	result := minMeetingRooms(intervals)
	fmt.Println(result) // Output: 2
}

// ✅ Time & Space Complexity
// Time: O(n log n) due to sorting and heap operations.
// Space: O(n) for the heap in the worst case.

// Would you like a version of this problem that returns the schedule of rooms used, 
// or a real-world extension like scheduling with room IDs?
