package main

import (
	"container/heap"
	"fmt"
)

// 🔢 Problem: Merge K Sorted Linked Lists
// You are given an array of k sorted linked lists.
// Merge them into one sorted linked list and return the head.

// ✅ Constraints:
// Total nodes: n
// Time complexity: O(n log k)
// Space: O(k) for heap

// ⚙️ Approach: Min-Heap (Priority Queue)
// Use a min-heap (priority queue) to always get the smallest head node among the lists.
// Push the first node of each list into the heap.
// Repeatedly:
// Pop the smallest node.
// Add it to the result list.
// Push its .Next into the heap (if it exists).
// Continue until the heap is empty.

// We’ll use Go’s container/heap package to implement the priority queue.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// PriorityQueue implements heap.Interface for ListNode pointers.
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}

	return dummy.Next
}

// Helper: Create a linked list from slice
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

// Helper: Print a linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := createList([]int{1, 4, 5})
	l2 := createList([]int{1, 3, 4})
	l3 := createList([]int{2, 6})
	lists := []*ListNode{l1, l2, l3}

	result := mergeKLists(lists)
	fmt.Print("Merged List: ")
	printList(result)
}

// 💡 Notes:
// Time complexity: O(n log k) where n = total number of nodes, k = number of lists.
// Space complexity: O(k) for the heap.

// Would you like to try a follow-up variation?
// For example: merging k lists without a heap (using divide-and-conquer)?
