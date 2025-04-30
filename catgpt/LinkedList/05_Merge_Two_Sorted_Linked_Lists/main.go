package main

import "fmt"

// 🔁 Problem: Merge Two Sorted Linked Lists
// Given two sorted linked lists, merge them into a single sorted linked list and
// return the head of the new list.

// Example:
// Input:
// List A: 1 → 3 → 5
// List B: 2 → 4 → 6
// Output:
// 1 → 2 → 3 → 4 → 5 → 6

// ✅ Approach:
// Use a dummy node to simplify handling the head.
// Use two pointers (l1 and l2) to traverse each list.
// Compare values, attach the smaller node to the result list.
// After one list runs out, attach the rest of the other list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Attach remaining list
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

// Helper to build a linked list from slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, num := range nums[1:] {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return head
}

// Helper to print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := createList([]int{1, 3, 5})
	l2 := createList([]int{2, 4, 6})

	fmt.Print("List 1: ")
	printList(l1)
	fmt.Print("List 2: ")
	printList(l2)

	merged := mergeTwoLists(l1, l2)
	fmt.Print("Merged List: ")
	printList(merged)
}

// This one is very popular in interviews (often a warm-up for merge-sort logic).
// Want to try:
// Recursive version of this?
// Or a harder variation, like merging k sorted linked lists (min-heap required)?