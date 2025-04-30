package main

import "fmt"

// Problem:
// Given two singly linked lists, return the node at which the two lists intersect.
// If the two linked lists have no intersection, return nil.
// Example:
// List A: 1 → 3 → 5 \
//                      → 7 → 9
// List B:       2 → 4 /

// Output: Node with value 7

// Constraints:
// No modification to the original lists.
// Must be O(n) time and O(1) space.

// Approach Explanation (Two Pointer Trick):
// Create two pointers pA and pB.
// Move each pointer one step at a time.
// When a pointer reaches the end of its list:
// Redirect it to the head of the other list.
// If there is an intersection, they will meet at the intersection node.
// If no intersection, both will reach nil at the same time.

// Why Does This Work?
// Suppose:
// Length of A = a + c
// Length of B = b + c (where c is the common part, and a and b are unique parts)
// After switching, both pointers traverse exactly a + b + c steps total.
// So they align perfectly at the intersection point.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

// Helper function to attach two lists at a given node
func attachLists(headA, headB *ListNode, attachNode *ListNode) {
	current := headB
	for current.Next != nil {
		current = current.Next
	}
	current.Next = attachNode
}

func main() {
	// Create two separate lists
	a := &ListNode{Val: 1}
	a.Next = &ListNode{Val: 3}
	a.Next.Next = &ListNode{Val: 5}

	b := &ListNode{Val: 2}
	b.Next = &ListNode{Val: 4}

	// Create the common intersection
	common := &ListNode{Val: 7}
	common.Next = &ListNode{Val: 9}

	// Attach the common part
	a.Next.Next.Next = common
	attachLists(a, b, common)

	// Find intersection
	intersection := getIntersectionNode(a, b)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.Val)
	} else {
		fmt.Println("No intersection")
	}
}

// ✅ Important Tip:
// This trick with "switching heads" is very elegant — it's O(n) time, O(1) space.
// Interviewers love when you solve this without extra memory like HashSets.

