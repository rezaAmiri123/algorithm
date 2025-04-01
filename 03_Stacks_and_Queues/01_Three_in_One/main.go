package main

import "errors"

// Three in One: Describe how you could use a single array to implement three stacks.
// SOLUTION
// Like many problems, this one somewhat depends on how well we'd like to support these stacks. If we're
// okay with simply allocating a fixed amount of space for each stack, we can do that. This may mean though
// that one stack runs out of space, while the others are nearly empty.
// Alternatively, we can be flexible in our space allocation, but this significantly increases the complexity of
// the problem.

// Approach 1: Fixed Division
// We can divide the array in three equal parts and allow the individual stack to grow in that limited space.
// Note: We will use the notation "[" to mean inclusive of an end point and "(" to mean exclusive of an end
// point.

// The code for this solution is below.
type FixedMultiStack struct {
	numbrOfStacks int
	stackCapacity int
	values        []int
	sizes         []int
}

var (
	ErrStackIsFull  = errors.New("stack is full")
	ErrStackIsEmpty = errors.New("stack is empty")
)

func newFixedMultiStack(numbrOfStacks, stackCapacity int) *FixedMultiStack {
	stack := &FixedMultiStack{
		numbrOfStacks: numbrOfStacks,
		stackCapacity: stackCapacity,
		values:        make([]int, stackCapacity),
		sizes:         make([]int, numbrOfStacks),
	}
	return stack
}

func (s *FixedMultiStack) push(stackNum, data int) error {
	// Check that we have space for the next element
	if s.isFull(stackNum) {
		return ErrStackIsFull
	}
	// Increment stack pointer and then update top value.
	s.sizes[stackNum]++
	s.values[stackNum] = data
	return nil
}

func (s *FixedMultiStack) pop(stackNum int) (int, error) {
	if s.isEmpty(stackNum) {
		return 0, ErrStackIsEmpty
	}
	topIndex :=s.indexOfTop(stackNum)
	value := s.values[topIndex]// get top
	s.values[topIndex]=0// clear
	s.sizes[stackNum]--// shrink
	return value,nil
}
func (s *FixedMultiStack)peek(stackNum int)(int,error){
	if s.isEmpty(stackNum){
		return 0, ErrStackIsEmpty
	}
	return s.values[s.indexOfTop(stackNum)],nil
}

func (s *FixedMultiStack) isFull(stackNum int) bool {
	return s.sizes[stackNum] == s.stackCapacity
}
func (s *FixedMultiStack) isEmpty(stackNum int) bool {
	return s.sizes[stackNum] == 0
}

// Returns index of the top of the stack.
func (s *FixedMultiStack) indexOfTop(stackNum int) int {
	offset := stackNum * s.stackCapacity
	size := s.sizes[stackNum]
	return offset + size - 1
}

// If we had additional information about the expected usages of the stacks, then we could modify this algo­
// rithm accordingly. For example, if we expected Stack 1 to have many more elements than Stack 2, we could
// allocate more space to Stack 1 and less space to Stack 2.

// to be continued at page 239
