package main

// Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks.
// SOLUTION
// Since the major difference between a queue and a stack is the order (first-in first-out vs. last-in first-out), we
// know that we need to modify peek() and pop() to go in reverse order. We can use our second stack to
// reverse the order of the elements (by popping sl and pushing the elements on to s2). In such an imple­
// mentation, on each peek() and pop() operation, we would pop everything from sl onto s2, perform
// the peek/pop operation, and then push everything back.
// This will work, but if two pop/peeks are performed back-to-back, we're needlessly moving elements. We
// can implement a "lazy" approach where we let the elements sit in s2 until we absolutely must reverse the
// elements.
// In this approach, stackNewest has the newest elements on top and stackOldest has the oldest
// elements on top. When we dequeue an element, we want to remove the oldest element first, and so we
// dequeue from stackOldest. If stackOldest is empty, then we want to transfer all elements from
// stackNewest into this stack in reverse order. To insert an element, we push onto stackNewest, since it
// has the newest elements on top.
// The code below implements this algorithm.
type Stack struct {
	array []int
	count int
}

func newStack() *Stack {
	return &Stack{
		array: make([]int, 0),
		count: -1,
	}
}
func (s *Stack) push(data int) {
	s.count++
	s.array = append(s.array, data)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		return -1000
	}
	data := s.array[s.count]
	s.count--
	return data
}

func (s *Stack) isEmpty() bool {
	if s.count <= 0 {
		return true
	}
	return false
}

func (s *Stack) peek() int {
	return s.array[s.count]
}

type Queue struct {
	stackNewest *Stack
	stackOldest *Stack
}

func NewQueue() *Queue {
	return &Queue{
		stackNewest: newStack(),
		stackOldest: newStack(),
	}
}

func (q *Queue) size() int {
	return q.stackNewest.count + q.stackOldest.count
}

func (q *Queue) add(data int) {
	// Push onto stackNewest, which always has the newest elements on top
	q.stackNewest.push(data)
}

// Move elements from stackNewest into stackOldest. This is usually done so that
// we can do operations on stackOldest.
func (q *Queue) shiftStacks() {
	if q.stackOldest.isEmpty() {
		for !q.stackNewest.isEmpty() {
			q.stackOldest.push(q.stackNewest.pop())
		}
	}
}

func (q *Queue) peek() int {
	q.shiftStacks()             //// Ensure stackOldest has the current elements
	return q.stackOldest.peek() // retrieve the oldest item.
}

func (q *Queue) remove() int {
	q.shiftStacks()            //// Ensure stackOldest has the current elements
	return q.stackOldest.pop() //pop the oldest item.
}

// During your actual interview, you may find that you forget the exact API calls. Don't stress too much if that
// happens to you. Most interviewers are okay with your asking for them to refresh your memory on little
// details. They're much more concerned with your big picture understanding.
