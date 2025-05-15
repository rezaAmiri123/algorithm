package main

import "fmt"

// Let’s walk through the “Merge K Sorted Lists” problem using divide and conquer,
// specifically without using a heap (i.e., no priority queue).
// This method is inspired by the merge step in Merge Sort.

// 🧠 Problem: Merge K Sorted Linked Lists
// Given an array of k linked lists, each linked list is sorted in ascending order.
// Goal: Merge all the linked lists into one sorted linked list and return it.

// 🔧 Constraints
// Do not use a heap.
// Use divide-and-conquer approach.

// ✅ Strategy: Divide and Conquer
// Pairwise merge the lists (merge two lists at a time).
// Recurse: Keep dividing the list of lists into halves, merge each half recursively.
// Base cases:
// If the list is empty, return nil.
// If it has one list, return it.
// Time Complexity:
// Each merge takes O(N) time where N is total number of nodes.
// Depth of recursion is log(k), and each level merges all elements.
// ⏱️ Overall: O(N log k), where N is total number of nodes across all lists.

// 🧪 Example
// lists = [
//   1 -> 4 -> 5,
//   1 -> 3 -> 4,
//   2 -> 6
// ]

// Output
// 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge two sorted linked lists
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}

// Divide and conquer to merge K lists
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeRange(lists, 0, len(lists)-1)
}

func mergeRange(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	l1 := mergeRange(lists, left, mid)
	l2 := mergeRange(lists, mid+1, right)
	return mergeTwoLists(l1, l2)
}

// Helper to build a linked list from a slice
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper to print a linked list
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " -> ")
		node = node.Next
	}
	fmt.Println("nil")
}

// Main function to test mergeKLists
func main() {
	// Create input: [[1,4,5],[1,3,4],[2,6]]
	lists := []*ListNode{
		buildList([]int{1, 4, 5}),
		buildList([]int{1, 3, 4}),
		buildList([]int{2, 6}),
	}

	// Merge all lists
	merged := mergeKLists(lists)

	// Print result
	printList(merged)
}