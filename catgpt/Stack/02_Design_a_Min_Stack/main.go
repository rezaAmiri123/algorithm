package main

import (
	"fmt"
	"math"
)

// 🔹 Question: Design a Min Stack
// Design a stack that supports these operations in O(1) time:
// Push(x int)
// Pop()
// Top() int
// GetMin() int → returns the minimum element in the stack

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64}, // sentinel value
	}
}

// Push adds to the stack and tracks the min
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	min := s.minStack[len(s.minStack)-1]
	if x < min {
		s.minStack = append(s.minStack, x)
	} else {
		s.minStack = append(s.minStack, min)
	}
}

// Pop removes from both stacks
func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

// Top returns the top of the stack
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

// GetMin returns the current minimum
func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

// Test
func main() {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // 0
	fmt.Println(minStack.GetMin()) // -2
}

// 🧠 How It Works
// stack: stores all values normally.
// minStack: stores the minimum at every level.
// When pushing: store the smaller of current min or new value.
// When popping: pop from both stacks.
// GetMin() just peeks the top of minStack.
// ⏱ All operations are O(1) time!

// Want to keep leveling up? I can give you:
// Evaluate Reverse Polish Notation
// Decode String with stack (like "3[a2[c]]")
// Next Greater Element
