package main

import (
	"fmt"
	"math/rand"
)

// Intersection: Given two (singly) linked lists, determine if the two lists intersect. Return the
// intersecting node. Note that the intersection is defined based on reference, not value. That is, if the
// kth node of the first linked list is the exact same node (by reference) as the jth node of the second
// linked list, then they are intersecting.

// Determining if there's an intersection.
// How would we detect if two linked lists intersect? One approach would be to use a hash table and just
// throw all the linked lists nodes into there. We would need to be careful to reference the linked lists by their
// memory location, not by their value.
// There's an easier way though. Observe that two intersecting linked lists will always have the same last node.
// Therefore, we can just traverse to the end of each linked list and compare the last nodes.
// How do we find where the intersection is, though

// Finding the intersecting node.
// One thought is that we could traverse backwards through each linked list. When the linked lists"split'; that's
// the intersection. Of course, you can't really traverse backwards through a singly linked list.
// If the linked lists were the same length, you could just traverse through them at the same time. When they
// collide, that's your intersection

// When they're not the same length, we'd like to just"chop off"-or ignore-those excess (gray) nodes.
// How can we do this? Well, if we know the lengths of the two linked lists, then the difference between those
// two linked lists will tell us how much to chop off.
// We can get the lengths at the same time as we get the tails of the linked lists (which we used in the first step
// to determine if there's an intersection).

// Putting it all together.
// We now have a multistep process.
// 1. Run through each linked list to get the lengths and the tails.
// 2. Compare the tails. If they are different (by reference, not by value), return immediately. There is no inter-
// section.
// 3. Set two pointers to the start of each linked list.
// 4. On the longer linked list, advance its pointer by the difference in lengths.
// 5. Now, traverse on each linked list until the pointers are the same.
// The implementation for this is below.
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

func findIntersection(list1, list2 *Node) *Node {
	if list1 == nil && list2 == nil {
		return nil
	}

	// Get tail and sizes.
	result1 := getTailAndSize(list1)
	result2 := getTailAndSize(list2)

	// If different tail nodes, then there's no intersection.
	if result1.tail != result2.tail {
		return nil
	}
	shorter := list1
	longer := list2
	if result1.size > result2.size {
		shorter = list2
		longer = list1
	}

	absSize := result1.size - result2.size
	if absSize < 0 {
		absSize = absSize * -1
	}
	// Advance the pointer for the longer linked list by difference in lengths.
	longer = getKthNode(longer, absSize)

	// Move both pointers until you have a collision.

	for shorter != longer {
		shorter = shorter.next
		longer = longer.next
	}
	// Return either one.
	return longer
}

func getKthNode(head *Node, k int) *Node {
	current := head
	for k > 0 && current != nil {
		current = current.next
		k--
	}
	return current
}

func getTailAndSize(list *Node) *Result {
	if list == nil {
		return nil
	}

	size := 1
	current := list
	for current != nil {
		size++
		current = current.next
	}
	return &Result{
		tail: current,
		size: size,
	}
}

type Result struct {
	tail *Node
	size int
}

func main() {
	common := &Node{data: 11}
	cur := common
	for i := 0; i < 5; i++ {
		cur.next = &Node{data: i}
		cur = cur.next
	}

	list1 := &Node{data: 23}
	cur = list1
	for i := 0; i < 8; i++ {
		cur.next = &Node{data: rand.Intn(10)}
		cur = cur.next
	}
	cur.next = common

	list2 := &Node{data: 23}
	cur = list2
	for i := 0; i < 4; i++ {
		cur.next = &Node{data: rand.Intn(10)}
		cur = cur.next
	}
	cur.next = common

	common.print()
	list1.print()
	list2.print()

	result := findIntersection(list1, list2)
	result.print()
}
