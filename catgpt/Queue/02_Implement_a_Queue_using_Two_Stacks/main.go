package main

import "fmt"

// 🧩 Problem: Implement a Queue using Two Stacks
// Description:
// Implement a queue with the following operations using two stacks:
// Push(x) – Push element x to the back of queue.
// Pop() – Removes the element from in front of queue.
// Peek() – Get the front element.
// Empty() – Return whether the queue is empty.

// ✍️ Solution in Go
// We’ll use two stacks:
// inStack – to handle push operations
// outStack – to handle pop/peek in FIFO order

// 🔍 How It Works:
// Push(x) → Just push to inStack.
// Pop() / Peek():
// If outStack is empty, we transfer all elements from inStack to outStack, reversing order.
// This makes the oldest element on top, simulating a queue.
// Empty() → Queue is empty only if both stacks are empty.

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) transfer() {
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			// Pop from inStack
			n := q.inStack[len(q.inStack)-1]
			q.inStack = q.inStack[:len(q.inStack)-1]
			// Push to outStack
			q.outStack = append(q.outStack, n)
		}
	}
}

func (q *MyQueue) Pop() int {
	q.transfer()
	front := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return front
}

func (q *MyQueue) Peek() int {
	q.transfer()
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())  // returns 1
	fmt.Println(queue.Pop())   // returns 1
	fmt.Println(queue.Empty()) // returns false
}
