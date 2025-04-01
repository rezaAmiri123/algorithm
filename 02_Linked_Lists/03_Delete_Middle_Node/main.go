package main

import "fmt"

// Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but
// the first and last node, not necessarily the exact middle) of a singly linked list, given only access to
// that node.
// EXAMPLE
// lnput:the node c from the linked list a->b->c->d->e->f
// Result: nothing is returned, but the new linked list looks like a->b->d- >e- >f

// SOLUTION
// In this problem, you are not given access to the head of the linked list. You only have access to that node.
// The solution is simply to copy the data from the next node over to the current node, and then to delete the
// next node.
// The code below implements this algorithm.
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
func deleteNode(n *Node) bool {
	if n == nil || n.next == nil {
		return false // failure
	}
	next := n.next
	n.data = next.data
	n.next = next.next
	return true
}

// Note that this problem cannot be solved if the node to be deleted is the last node in the linked list. That's
// okay-your interviewer wants you to point that out, and to discuss how to handle this case. You could, for
// example, consider marking the node as dummy.

func main() {
	head := &Node{data: 10}
	cur := head
	for i := 0; i < 10; i++ {
		cur.next = &Node{data: i}
		cur = cur.next
	}
	head.print()
	deleteNode(head.next.next.next)
	head.print()
}
