package main

import "fmt"

// Route Between Nodes: Given a directed graph, design an algorithm to find out whether there is a
// route between two nodes.

// page 120, 252
// SOLUTION
// --------------------- -------
// This problem can be solved by just simple graph traversal, such as depth-first search or breadth-first search.
// We start with one of the two nodes and, during traversal, check if the other node is found. We should mark
// any node found in the course of the algorithm as "already visited" to avoid cycles and repetition of the
// nodes.
// The code below provides an iterative implementation of breadth-first search.
type State string

var (
	Unvisted State = "unvisited"
	Visted   State = "visited"
	Visting  State = "visiting"
)

type Node struct {
	name     string
	children []*Node
	state    State
	next     *Node
}

type Tree struct {
	root *Node
}

type Graph struct {
	nodes []*Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}
func (s *LinkedList) add(n *Node) {
	s.size++
	if s.head == nil && s.tail == nil {
		s.head = n
		s.tail = n
		return
	}
	s.tail.next = n
	s.tail = n
}
func (s *LinkedList) removeFirst() *Node {
	s.size--
	if s.tail == s.head{
		n := s.head
		s.head =nil
		s.tail = nil
		return n
	}
	n := s.head
	s.head = s.head.next
	return n
}

func (s *LinkedList) isEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}
func (s *LinkedList) print() {
	cur := s.tail
	for cur != nil {
		fmt.Printf("%s-->", cur.name)
		cur = cur.next
	}
	fmt.Println()
}
func search(g *Graph, start, end *Node) bool {
	if start == end {
		return true
	}

	q := newLinkedList()
	q.print()
	// operates as Queue
	for _, u := range g.nodes {
		u.state = Unvisted
	}

	start.state = Visted
	q.add(start)
	var u *Node
	for !q.isEmpty() {
		u = q.removeFirst() // i.e., dequeue()
		if u != nil {
			fmt.Println("len(u.children): ",len(u.children), u.name)
			for _, v := range u.children {
				if v.state == Unvisted {
					if v == end {
						return true
					} else {
						v.state = Visting
						q.add(v)
						fmt.Println("v: ", v.name)
					}
				}
			}
			u.state = Visted
		}
	}
	return false
}
// It may be worth discussing with your interviewer the tradeoffs between breadth-first search and depth-first
// search for this and other problems. For example, depth-first search is a bit simpler to implement since it can
// be done with simple recursion. Breadth-first search can also be useful to find the shortest path, whereas
// depth-first search may traverse one adjacent node very deeply before ever going onto the immediate
// neighbors.

func main() {
	nodes := []*Node{}
	for i := 0; i < 10; i++ {
		nodes = append(nodes, &Node{name: fmt.Sprintf("%d", i)})
	}
	nodes[5].children = nodes[:5]
	nodes[4].children = nodes[5:]
	nodes[0].children = nodes[1:8]

	g := &Graph{nodes: nodes}

	fmt.Println(search(g, nodes[0], nodes[9]))

}
