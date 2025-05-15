package main

import (
	"fmt"
	"sort"
)

// ✅ Problem: Meeting Rooms (I)
// Description:
// Given an array of meeting time intervals consisting of start and end times
// [[s1,e1],[s2,e2],...], determine if a person could attend all meetings
// (i.e., none of the intervals overlap).

// 🔍 Example
// Input:
// intervals := [][]int{{0,30},{5,10},{15,20}}
// Output:
// false

// Explanation:
// The meetings [0,30] and [5,10] overlap.
// Therefore, it is not possible to attend all meetings.

func canAttendMeetings(intervals [][]int) bool {
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Check for overlapping meetings
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	result := canAttendMeetings(intervals)
	fmt.Println(result) // Output: false
}

// 🧠 Explanation
// We sort the meetings by start time.
// Then we loop through the intervals and check if the start time of a meeting is 
// before the end time of the previous one.
// If it is, there's an overlap → return false.

// Would you like to try the harder version next: 
// Meeting Rooms II (where you compute the minimum number of meeting rooms needed)?
