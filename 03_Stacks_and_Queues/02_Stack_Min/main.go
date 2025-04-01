package main

import (
	"errors"
	"fmt"
)

// Stack Min: How would you design a stack which, in addition to push and pop, has a function min
// which returns the minimum element? Push, pop and min should all operate in 0(1) time.

// SOLUTION
// The thing with minimums is that they don't change very often. They only change when a smaller element
// is added.
// One solution is to have just a single int value, minValue, that's a member of the Stack class. When
// minValue is popped from the stack, we search through the stack to find the new minimum. Unfortunately,
// this would break the constraint that push and pop operate in 0(1) time.
// To further understand this question, let's walk through it with a short example:
// push(5); // stack is {5}, min is 5
// push(6); // stack is {6, 5}, min is 5
// push(3); // stack is {3, 6, 5}, min is 3
// push(7); // stack is {7, 3, 6, 5}, min is 3
// pop(); // pops 7. stack is {3, 6, 5}, min is 3
// pop(); // pops 3. stack is {6, 5}. min is 5.

// Observe how once the stack goes back to a prior state ({ 6, 5} ), the minimum also goes back to its prior
// state (5). This leads us to our second solution.
// If we kept track of the minimum at each state, we would be able to easily know the minimum. We can do
// this by having each node record what the minimum beneath itself is. Then, to find the min, you just look at
// what the top element thinks is the min.
// When you push an element onto the stack, the element is given the current minimum. It sets its "local
// min"to be the min.

type Node struct {
	next *Node
	data int
	min  int
}

func (n *Node) print() {
	cur := n
	for cur != nil {
		fmt.Printf("(%d, %d)-->", cur.data, cur.min)
		cur = cur.next
	}
	fmt.Println()
}

type Stack struct {
	node  *Node
	count int
}

func (s *Stack) print() {
	s.node.print()
}

func (s *Stack) push(data int) {
	node := &Node{data: data}
	if s.node == nil {
		s.node = node
		s.node.min = data
		return
	}
	if data > s.node.min {
		node.min = s.node.min
	} else {
		node.min = data
	}
	node.next = s.node
	s.node = node
	s.count++
}
func (s *Stack) pop() (int, error) {
	if s.count == 0 {
		return 0, errors.New("stack empty")
	}
	s.count--
	data := s.node.data
	s.node = s.node.next
	return data, nil
}

func (s *Stack) min() (int, error) {
	if s.count == 0 {
		return 0, errors.New("stack empty")
	}
	return s.node.min, nil
}

// There's just one issue with this: if we have a large stack, we waste a lot of space by keeping track of the min
// for every single element. Can we do better?
// We can (maybe) do a bit better than this by using an additional stack which keeps track of the mins.
type StackImproved struct {
	node    *Node
	count   int
	minNode *Node
}

func (s *StackImproved) print() {
	s.node.print()
}
func (s *StackImproved) printMin() {
	s.minNode.print()
}

func (s *StackImproved) push(data int) {
	node := &Node{data: data}
	if s.node == nil {
		s.node = node
		s.minNode = node
		return
	}
	if data < s.minNode.data {
		cur := &Node{data: data, next: s.minNode}
		s.minNode = cur
	}
	node.next = s.node
	s.node = node
	s.count++
}
func (s *StackImproved) pop() (int, error) {
	if s.count == 0 {
		return 0, errors.New("stack empty")
	}
	s.count--
	data := s.node.data
	s.node = s.node.next
	if data == s.minNode.data {
		s.minNode = s.minNode.next
	}
	return data, nil
}

func (s *StackImproved) min() (int, error) {
	if s.count == 0 {
		return 0, errors.New("stack empty")
	}
	return s.minNode.data, nil
}

func main() {
	s := &Stack{}
	s2 := &StackImproved{}
	for i := 10; i > 0; i-- {
		s.push(i)
		s2.push(i)
	}
	s.print()
	s2.print()
	s2.printMin()
}
