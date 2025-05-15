package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// ✅ Problem: Assign Meetings to Rooms with IDs
// Description:
// You're given an array of meeting intervals [[start1,end1], [start2,end2], ...].
// Assign each meeting to a room such that no two meetings in the same room overlap.
// Return a list of meeting assignments where each entry includes the meeting interval
// and the room ID it was assigned to.

// 🔍 Example
// Input:
// intervals := [][]int{{0,30},{5,10},{15,20},{35,40}}
// Output (room IDs may vary, but overlap is avoided):

// [
//   {[0 30] Room 0},
//   {[5 10] Room 1},
//   {[15 20] Room 1},
//   {[35 40] Room 0}
// ]

// Explanation:
// [0,30] uses Room 0.
// [5,10] conflicts with it, so it goes to Room 1.
// [15,20] can reuse Room 1.
// [35,40] goes back to Room 0 (now free).

// Struct to hold meeting with its original index for result mapping
type Meeting struct {
	start int
	end   int
	index int
}

// Room struct to manage heap with room ID and available time
type Room struct {
	endTime int
	id      int
}

// Priority queue implementation (min-heap) by end time
type RoomHeap []Room

func (h RoomHeap) Len() int           { return len(h) }
func (h RoomHeap) Less(i, j int) bool { return h[i].endTime < h[j].endTime }
func (h RoomHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RoomHeap) Push(x interface{}) {
	*h = append(*h, x.(Room))
}

func (h *RoomHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func assignMeetingRooms(intervals [][]int) []string {
	if len(intervals) == 0 {
		return []string{}
	}

	// Prepare meeting list with index tracking
	meetings := make([]Meeting, len(intervals))
	for i, interval := range intervals {
		meetings[i] = Meeting{interval[0], interval[1], i}
	}

	// Sort by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].start < meetings[j].start
	})

	h := &RoomHeap{}
	heap.Init(h)

	roomCount := 0
	assignments := make([]string, len(intervals))

	for _, m := range meetings {
		// Reuse room if possible
		if h.Len() > 0 && (*h)[0].endTime <= m.start {
			room := heap.Pop(h).(Room)
			assignments[m.index] = fmt.Sprintf("{[%d %d] Room %d}", m.start, m.end, room.id)
			room.endTime = m.end
			heap.Push(h, room)
		} else {
			// Assign new room
			assignments[m.index] = fmt.Sprintf("{[%d %d] Room %d}", m.start, m.end, roomCount)
			heap.Push(h, Room{endTime: m.end, id: roomCount})
			roomCount++
		}
	}

	return assignments
}

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}, {35, 40}}
	result := assignMeetingRooms(intervals)
	for _, r := range result {
		fmt.Println(r)
	}
}

// 🧠 Why This Matters
// This version reflects real-world scheduling:
// Room availability based on timing.
// Dynamic room allocation.
// Optimization of resources (minimal room usage).
// Would you like to extend this even further to prioritize room preferences 
// or include room capacity constraints?
