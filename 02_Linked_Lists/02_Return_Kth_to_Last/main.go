package main

import (
	"fmt"
)

// Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
// We will approach this problem both recursively and non-recursively. Remember that recursive solutions are
// often cleaner but less optimal. For example, in this problem, the recursive implementation is about half the
// length of the iterative solution but also takes 0( n) space, where n is the number of elements in the linked
// list.
// Note that for this solution, we have definedk such that passing ink = 1 would return the last element,k
// = 2 would return to the second to last element, and so on. It is equally acceptable to definek such thatk
// = 0 would return the last element.

// Solution #1: If linked list size is known
// If the size of the linked list is known, then thekth to last element is the ( length - k)th element. We can
// just iterate through the linked list to find this element. Because this solution is so trivial, we can almost be
// sure that this is not what the interviewer intended.

// Solution #2: Recursive
// This algorithm recurses through the linked list. When it hits the end, the method passes back a counter set
// to 0. Each parent call adds 1 to this counter. When the counter equalsk, we know we have reached thekth
// to last element of the linked list.
// Implementing this is short and sweet-provided we have a way of"passing back" an integer value through
// the stack. Unfortunately, we can't pass back a node and a counter using normal return statements. So how
// do we handle this?
// Approach A: Don't Return the Element.
// One way to do this is to change the problem to simply printing thekth to last element. Then, we can pass
// back the value of the counter simply through return values.

type Node struct {
	next  *Node
	value int
}

func (n *Node) print() {
	cur := n
	for cur != nil {
		fmt.Printf("%d-->", cur.value)
		cur = cur.next
	}
	fmt.Println()
}
func printKthToLast(head *Node, k int) int {
	if head == nil {
		return 0
	}
	index := printKthToLast(head.next, k) + 1
	if index == k {
		fmt.Printf("%d the to last node is %d\n", k, head.value)
	}
	return index
}

// Of course, this is only a valid solution if the interviewer says it is valid

// A second way to solve this is to use C++ and to pass values by reference. This allows us to return the node
// value, but also update the counter by passing a pointer to it.
func nthToLast(head *Node, k int, i *int) *Node {
	if head == nil {
		return nil
	}
	nd := nthToLast(head.next, k, i)
	*i = *i + 1
	if *i == k {
		return head
	}
	return nd
}

// Approach C: Create a Wrapper Class.
// We described earlier that the issue was that we couldn't simultaneously return a counter and an index. If
// we wrap the counter value with simple class (or even a single element array), we can mimic passing by
// reference
// at page 221

// Solution #3: Iterative
// A more optimal, but less straightforward, solution is to implement this iteratively. We can use two pointers,
// pl and p2. We place them k nodes apart in the linked list by putting p2 at the beginning and moving pl
// k nodes into the list. Then, when we move them at the same pace, pl will hit the end of the linked list after
// LENGTH - k steps. At that point, p2 will be LENGTH - k nodes into the list, or k nodes from the end.
// The code below implements this algorithm.
func nthToLastItrator(head *Node, k int) *Node {
	p1 := head
	p2 := head
	// Move pl k nodes into the list.
	for i := 0; i < k; i++ {
		// Out of bounds
		if p1 == nil {
			return nil
		}
		p1 = p1.next
	}
	// Move them at the same pace. When pl hits the end, p2 will be at the right element.
	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
}

// This algorithm takes O(n) time and 0(1) space.

func main() {
	n := &Node{value: 10}
	cur := n
	for i := 0; i < 10; i++ {
		cur.next = &Node{value: i}
		cur = cur.next
	}
	n.print()
	fmt.Println("printKthToLast")
	printKthToLast(n, 4)
	i := 0
	nth := nthToLast(n, 4, &i)
	fmt.Println("nthToLast: ", nth.value)

	nth = nthToLastItrator(n, 4)
	fmt.Println("nthToLastItrator: ", nth.value)

}
