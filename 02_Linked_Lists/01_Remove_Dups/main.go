package main

import "fmt"

// Remove Dups: Write code to remove duplicates from an unsorted linked list.
// FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?
// SOLUTION
// In order to remove duplicates from a linked list, we need to be able to track duplicates. A simple hash table
// will work well here.
// In the below solution, we simply iterate through the linked list, adding each element to a hash table. When
// we discover a duplicate element, we remove the element and continue iterating. We can do this all in one
// pass since we are using a linked list.

type Node struct {
	next *Node
	data int
}

func (n *Node) AppendToTail(d int) {
	for n.next != nil {
		n = n.next
	}
	n.next = &Node{data: d}
}

func (n *Node) Print() {
	for n != nil {
		fmt.Print(n.data, "-->")
		n = n.next
	}
	fmt.Println()
}

func deleteDups(n *Node) {
	set := make(map[int]struct{})
	var pervious *Node
	for n != nil {
		_, exists := set[n.data]
		if exists {
			pervious.next = n.next
		} else {
			set[n.data] = struct{}{}
			pervious = n
		}
		n = n.next
	}
}

// The above solution takes O(N) time, where N is the number of elements in the linked list.
// Follow Up: No Buffer Allowed
// lfwe don't have a buffer, we can iterate with two pointers: current which iterates through the linked list,
// and runner which checks all subsequent nodes for duplicates.
func deleteDupsNoBuffer(head *Node){
	current := head
	for current != nil{
		// Remove all future nodes that have the same value
		runner := current
		for runner.next!= nil{
			if runner.next.data == current.data{
				runner.next = runner.next.next
			}else{
				runner = runner.next
			}
		}
		current = current.next
	}
}
// This code runs in O ( 1) space, but O ( N2) time.

func main() {
	n := &Node{data: 10}
	for i := 0; i < 10; i++ {
		n.AppendToTail(i)
	}
	n.AppendToTail(6)
	n.Print()
	deleteDups(n)
	n.Print()
	n.AppendToTail(6)
	n.Print()
	deleteDupsNoBuffer(n)
	n.Print()
}
