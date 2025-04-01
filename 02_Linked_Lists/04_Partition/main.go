package main

import (
	"fmt"
	"math/rand"
)

// Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
// before all nodes greater than or equal to x. If x is contained within the list the values of x only need
// to be after the elements less than x (see below). The partition element x can appear anywhere in the
// "right partition"; it does not need to appear between the left and right partitions.
// EXAMPLE
// Input:3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition= 5]
// Output:3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8

// SOLUTION
// If this were an array, we would need to be careful about how we shifted elements. Array shifts are very
// expensive.
// However, in a linked list, the situation is much easier. Rather than shifting and swapping elements, we can
// actually create two different linked lists: one for elements less than x, and one for elements greater than or
// equal to x.
// We iterate through the linked list, inserting elements into our before list or our after list. Once we reach
// the end of the linked list and have completed this splitting, we merge the two lists.
// This approach is mostly "stable" in that elements stay in their original order, other than the necessary move­
// ment around the partition. The code below implements this approach.
type Node struct {
	next *Node
	data int
}

func (n *Node) print() {
	for n != nil {
		fmt.Printf("%d-->", n.data)
		n = n.next
	}
	fmt.Println()
}

func partition(head *Node, x int) *Node {
	var beforeStart, afterStart, beforeEnd, afterEnd *Node
	// Partition list
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = nil
		if cur.data < x {
			// Insert node into end of before list
			if beforeStart == nil {
				beforeStart = cur
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = cur
				beforeEnd = cur
			}
		} else {
			// Insert node into end of after list
			if afterStart == nil {
				afterStart = cur
				afterEnd = afterStart
			} else {
				afterEnd.next = cur
				afterEnd = cur
			}
		}
		cur = next
	}
	if beforeStart == nil {
		return afterStart
	}

	// Merge before list and after list
	beforeEnd.next = afterStart
	return beforeStart
}

// If it bugs you to keep around four different variables for tracking two linked lists, you're not alone. We can
// make this code a bit shorter.
// If we don't care about making the elements of the list "stable" (which there's no obligation to, since the
// interviewer hasn't specified that), then we can instead rearrange the elements by growing the list at the
// head and tail.
// In this approach, we start a"new" list (using the existing nodes). Elements bigger than the pivot element are
// put at the tail and elements smaller are put at the head. Each time we insert an element, we update either
// the head or tail.
func partitionImproved(node *Node, x int) *Node {
	head := node
	tail := node

	for node != nil {
		next := node.next
		if node.data < x {
			// Insert node at head.
			node.next = head
			head = node
		} else {
			// Insert node at tail.
			tail.next = node
			tail = node
		}
		node = next
	}
	tail.next = nil
	// The head has changed, so we need to return it to the user.
	return head
}

// There are many equally optimal solutions to this problem. If you came up with a different one, that's okay!

func main() {
	head := &Node{data: 10}
	cur := head
	for i := 0; i < 10; i++ {
		cur.next = &Node{data: rand.Intn(20)}
		cur = cur.next
	}
	head.print()
	p := partition(head, 10)
	p.print()
	fmt.Println("improved")
	cur = head
	for i := 0; i < 10; i++ {
		cur.next = &Node{data: rand.Intn(20)}
		cur = cur.next
	}
	head.print()
	p = partitionImproved(head, 10)
	p.print()

}
