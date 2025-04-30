package main

import "fmt"

// Problem: Reverse a Singly Linked List
// Given the head of a singly linked list, reverse the list, and return the new head.

// Example
// Input:
// 1 -> 2 -> 3 -> 4 -> 5 -> nil
// Output:
// 5 -> 4 -> 3 -> 2 -> 1 -> nil

// Definition for singly-linked list.
type ListNode struct{
	Val int
	Next *ListNode
}

// Function to reverse the linked list
func reverseList(head *ListNode)*ListNode{
	var prev *ListNode = nil
	current := head

	for current!= nil{
		nextTemp := current.Next // store next node
        current.Next = prev       // reverse current node's pointer
        prev = current            // move prev forward
        current = nextTemp        // move current forward
	}

	return prev
}

// Helper function to print the linked list
func printList(head *ListNode){
	for head != nil{
		fmt.Print(head.Val, " -> ")
		head=head.Next
	}
	fmt.Println("nil")
}


// Main function to run the example
func main() {
    // Create the list 1 -> 2 -> 3 -> 4 -> 5 -> nil
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 5}

    fmt.Println("Original List:")
    printList(head)

    // Reverse the list
    reversedHead := reverseList(head)

    fmt.Println("Reversed List:")
    printList(reversedHead)
}

// Explanation
// We iterate through the linked list.
// We keep track of:
// prev (initially nil) → the previous node.
// current → the current node we are visiting.
// At each step:
// Save current.Next in nextTemp.
// Change current.Next to prev (this reverses the link).
// Move prev and current one step forward.
// At the end, prev will be the new head of the reversed list.

// Would you like me to give you a follow-up harder problem based on Linked Lists too? 
// 🚀 (Like detecting a cycle, merging two sorted lists, etc.)
