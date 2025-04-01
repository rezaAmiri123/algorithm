package main

import (
	"fmt"
	"math/rand"
)

// Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use
// an additional temporary stack, but you may not copy the elements into any other data structure
// (such as an array). The stack supports the following operations: push, pop, peek, and is Empty.

// SOLUTION
// One approach is to implement a rudimentary sorting algorithm. We search through the entire stack to find
// the minimum element and then push that onto a new stack. Then, we find the new minimum element
// and push that. This will actually require a total of three stacks: s 1 is the original stack, s2 is the final sorted
// stack, and s3 acts as a buffer during our searching of sl. To search sl for each minimum, we need to pop
// elements from sl and push them onto the buffer, s3.
// Unfortunately, this requires two additional stacks, and we can only use one. Can we do better? Yes.
// Rather than searching for the minimum repeatedly, we can sort sl by inserting each element from sl in
// order into s2. How would this work?
// Imagine we have the following stacks, where s2 is"sorted" and sl is not:
type Stack struct {
	array []int
	size  int
}

func newStack() *Stack {
	return &Stack{
		array: make([]int, 0),
		size:  -1,
	}
}
func (s *Stack) push(data int) {
	s.size++
	if len(s.array) > s.size+1 {
		s.array[s.size] = data
	}
	s.array = append(s.array, data)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("stack is empty")
	}
	data := s.array[s.size]
	s.size--
	return data
}
func (s *Stack) peek() int {
	if s.isEmpty() {
		fmt.Println("stack is empty")
	}
	return s.array[s.size]
}

func (s *Stack) isEmpty() bool {
	if s.size <= 0 {
		return true
	}
	return false
}

func sort(s *Stack) {
	r := newStack()
	for !s.isEmpty() {
		// Insert each element in s in sorted order into r.
		temp := s.pop()
		for !r.isEmpty() && r.peek() > temp {
			s.push(r.pop())
		}
		r.push(temp)
	}
	// Copy the elements from r back into s
	for !r.isEmpty() {
		s.push(r.pop())
	}
}
// This algorithm is O ( N2 ) time and O ( N) space.
// If we were allowed to use unlimited stacks, we could implement a modified quicksort or mergesort.
// With the mergesort solution, we would create two extra stacks and divide the stack into two parts. We
// would recursively sort each stack, and then merge them back together in sorted order into the original
// stack. Note that this would require the creation of two additional stacks per level of recursion.
// With the quicksort solution, we would create two additional stacks and divide the stack into the two stacks
// based on a pivot element. The two stacks would be recursively sorted, and then merged back together
// into the original stack. Like the earlier solution, this one involves creating two additional stacks per level of
// recursion.

func main() {
	s := newStack()
	for i := 0; i < 10; i++ {
		s.push(rand.Intn(10))
	}
	fmt.Println(s.array)
	sort(s)
	fmt.Println(s.array[:10])
}
