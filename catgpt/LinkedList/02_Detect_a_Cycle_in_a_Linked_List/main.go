package main

import "fmt"

// Problem: Detect a Cycle in a Linked List
// Given the head of a linked list, determine if the linked list has a cycle in it.
// A cycle means a node's next pointer points back to a previous node.

// Example
// Input:
// 3 -> 2 -> 0 -> -4 -> (points back to 2)
// Output:
// true (because there is a cycle)
// Input:
// 1 -> 2 -> nil
// Output:
// false (no cycle)

// Golang Solution (Floyd’s Tortoise and Hare Algorithm)
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to detect a cycle
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// Helper function to create a list with cycle
func createCyclicList() *ListNode {
	head := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // cycle here

	return head
}

// Main function to test
func main() {
	head := createCyclicList()

	if hasCycle(head) {
		fmt.Println("Cycle detected!")
	} else {
		fmt.Println("No cycle detected.")
	}
}

// Explanation
// We use two pointers:
// Slow Pointer → moves one step at a time.
// Fast Pointer → moves two steps at a time.
// Idea:
// If there is a cycle, fast and slow pointers will eventually meet.
// If there is no cycle, fast will reach the end (nil) first.

// At some point, slow == fast → cycle detected!
// Would you also like me to show you a small variation:
// "Find where the cycle starts"?
// (it’s a common follow-up question!) 🚀
// Want to try it too? 🎯