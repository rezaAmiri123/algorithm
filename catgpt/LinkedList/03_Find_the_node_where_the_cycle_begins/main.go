package main

import "fmt"

// Problem:
// Find the node where the cycle begins in a linked list.
// If there is no cycle, return nil.

// Example:
// Input: 1 -> 2 -> 3 -> 4 -> 5 -> 2 (cycle starts at node 2)
// Output: Node with value 2

// Approach Explanation (Floyd's Cycle Detection + Math Trick):
// First, detect if there is a cycle using slow and fast.
// Once they meet, do this:
// Keep slow at the meeting point.
// Move a new pointer (start) from the head.
// Move both start and slow one step at a time.
// The node where they meet again is the start of the cycle.

// Why Does This Work?
// (Quick intuition):
// Let's say:
// L1 = distance from head to start of cycle
// L2 = distance inside the cycle to meeting point
// C = length of the full cycle
// Math shows that:
// Distance from meeting point to cycle start is exactly L1 steps.
// So, moving two pointers one-by-one, they'll meet exactly at the start of the cycle.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			start := head
			for start != slow {
				start = start.Next
				slow = slow.Next
			}
			return start
		}
	}

	return nil
}

// Helper functions
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func createCycle(head *ListNode, pos int) {
	if head == nil || pos < 0 {
		return
	}
	cycleNode := head
	for i := 0; i < pos; i++ {
		cycleNode = cycleNode.Next
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = cycleNode
}

func main() {
	list := createList([]int{3, 2, 0, -4})
	createCycle(list, 1) // cycle at node with value 2
	entry := detectCycle(list)
	if entry != nil {
		fmt.Println("Cycle starts at node with value:", entry.Val)
	} else {
		fmt.Println("No cycle detected")
	}
}

// ✅ Important Tip:
// Many candidates know how to detect a cycle, but few know how to find where it starts.
// If you can explain it + code it cleanly, you look very strong.
// Would you like one even cooler linked list challenge after this? 🤔
// Something like:
// Copy a linked list with random pointers
// or
// Find the intersection point of two linked lists
// Both are classic advanced questions! 🚀
