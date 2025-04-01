package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Stack of Plates: Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
// Therefore, in real life, we would likely start a new stack when the previous stack exceeds some
// threshold. Implement a data structure SetOfStacks that mimics this. SetOfStacks should be
// composed of several stacks and should create a new stack once the previous one exceeds capacity.
// SetOfStacks.push() and SetOfStacks. pop() should behave identically to a single stack
// (that is, pop() should return the same values as it would if there were just a single stack).
// FOLLOW UP
// Implement a function popAt(int index) which performs a pop operation on a specific sub­
// stack.

// SOLUTION
// In this problem, we've been told what our data structure should look like:

type Stack struct {
	array    [][]int
	capacity int
	stackID  int
	objID    int
}

// know that push () should behave identically to a single stack, which means that we need push () to
// call push () on the last stack in the array of stacks. We have to be a bit careful here though: if the last stack
// is at capacity, we need to create a new stack. Our code should look something like this:
func (s *Stack) push(data int) {
	if s.objID+1 >= s.capacity {
		s.stackID++
		s.objID = 0
		s.addStack()
	}else{
		s.objID++
	}
	s.array[s.stackID][s.objID] = data
}
func(s *Stack)addStack(){
	s.array = append(s.array, make([]int, s.capacity))
}

func (s *Stack) pop() (int, error) {
	if s.stackID == 0 && s.objID == 0 {
		return 0, errors.New("stack empty")
	}
	if s.objID == 0 {
		s.array[s.stackID]=nil
		s.stackID--
		s.objID = s.capacity - 1
	}
	data:= s.array[s.stackID][s.objID]
	s.objID--
	return data,nil
}

// Follow Up: Implement popAt(int index)
// to be continued at page 245

func main(){
	s := &Stack{
		array: [][]int{},
		capacity: 5,
		objID: -1,
	}
	s.addStack()
	for i:=0;i<10;i++{
		s.push(rand.Intn(10))
	}
	fmt.Println(s.array)
	for i:=0;i<5;i++{
		s.pop()
	}
	fmt.Println(s.array)
	fmt.Println(s.stackID, "  ", s.objID)
}