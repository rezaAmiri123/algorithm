package main

import "fmt"

// 🔄 Problem: Sliding Window Maximum
// Description:
// Given an array of integers nums and an integer k, there is a sliding window of size k moving from the very left to the very right of the array. You need to return the maximum value in each window.

// ✅ Solution with Explanation (Using Deque for Optimization)
// We’ll use a deque to store the indices of elements in the current window. The front of the deque will always have the index of the max element for the current window.

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var result []int
	var deque []int // stores indices

	for i := 0; i < len(nums); i++ {
		// Remove elements outside the current window from front
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// Remove smaller elements in deque from the back
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Add current index
		deque = append(deque, i)

		// Add to result once window hits size k
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k)) // Output: [3 3 5 5 6 7]
}

// 🧠 Explanation:
// Deque stores indices, not values.
// We remove indices from the front if they’re out of window.
// We remove from back if current element is greater than the one at those indices 
// (no point keeping smaller values).
// The front of the deque always holds the maximum for the current window.
// Want to try a simpler one first, like implementing a queue from scratch 
// or using two stacks? Or ready to move to a follow-up like sliding window minimum?