package main

// Sure! Practicing the Median in a Stream problem using two heaps is a great way to prepare 
// for algorithm interviews. Let's go over:
// Problem Statement (Medium level)
// Approach using two heaps
// Example Input and Step-by-Step Explanation
// Go (Golang) Implementation

// 💡 Problem Statement (Medium)
// Design a data structure that supports the following operations:
// AddNum(int num): Add a number to the data stream.
// FindMedian() float64: Return the median of all elements so far.

// Constraints:
// Each operation should be optimized for time (logarithmic or constant if possible).
// You will be given a sequence of numbers, one at a time.

// ⚙️ Approach: Two Heaps
// Use two heaps to maintain the lower and upper halves of the numbers:
// maxHeap: A max heap for the lower half (we use a min-heap with negated numbers in Go).
// minHeap: A min heap for the upper half.

// Balance rule:
// Keep sizes of both heaps almost equal (the difference can be at most 1).
// Ensure maxHeap always has the smaller half.

// Median logic:
// If both heaps are equal in size → median = average of tops
// If maxHeap has more → median = top of maxHeap

// 🧪 Example Input
// AddNum(5)
// AddNum(15)
// AddNum(1)
// AddNum(3)
// FindMedian()

// Step-by-step:
// Add 5 → maxHeap: [5], minHeap: []
// Add 15 → maxHeap: [5], minHeap: [15]
// Add 1 → rebalance → maxHeap: [5, 1], minHeap: [15]
// Add 3 → rebalance → maxHeap: [3, 1], minHeap: [5, 15]
// Median: average of 3 and 5 → (3 + 5) / 2 = 4.0

import (
	"container/heap"
	"fmt"
)

// MaxHeap (invert values)
type MaxHeap []int
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeap (regular)
type MinHeap []int
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MedianFinder struct
type MedianFinder struct{
	lower *MaxHeap
	upper *MinHeap
}

func Constructor()MedianFinder{
	l := &MaxHeap{}
	u:=&MinHeap{}
	heap.Init(l)
	heap.Init(u)
	return MedianFinder{lower: l,upper: u}
}

func (mf *MedianFinder)AddNum(num int){
	heap.Push(mf.lower,num)

	// Balance: move from lower to upper
	heap.Push(mf.upper,heap.Pop(mf.lower))

	// Rebalance if needed
	if mf.upper.Len()>mf.lower.Len(){
		heap.Push(mf.lower,heap.Pop(mf.upper))
	}
}

func (mf *MedianFinder)FindMedian()float64{
	if mf.lower.Len()>mf.upper.Len(){
		return float64((*mf.lower)[0])
	}
	return float64((*mf.lower)[0]+(*mf.upper)[0])/2.0
}

func main() {
	mf := Constructor()
	mf.AddNum(5)
	mf.AddNum(15)
	mf.AddNum(1)
	mf.AddNum(3)
	fmt.Println("Median:", mf.FindMedian()) // Output: 4.0
}
