package main

import "fmt"

// Palindrome: Implement a function to check if a linked list is a palindrome.

// SOLUTION
// To approach this problem, we can picture a palindrome like 0 - > 1 - > 2 - > 1 - > 0. We know that,
// since it's a palindrome, the list must be the same backwards and forwards. This leads us to our first solution
// Solution #1: Reverse and Compare
// Our first solution is to reverse the linked list and compare the reversed list to the original list. If they're the
// same, the lists are identical.
// Note that when we compare the linked list to the reversed list, we only actually need to compare the first
// half of the list. If the first half of the normal list matches the first half of the reversed list, then the second half
// of the normal list must match the second half of the reversed list.
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

func isPalindrome(head *Node) bool {
	reversed := reverseAndClone(head)
	return isEqual(head, reversed)
}

func reverseAndClone(node *Node) *Node {
	var head *Node
	for node != nil {
		n := &Node{next: head, data: node.data}
		head = n
		node = node.next
	}
	return head
}

func isEqual(one, two *Node) bool {
	for one != nil && two != nil {
		if one.data != two.data {
			return false
		}
		one = one.next
		two = two.next
	}
	return one == nil && two == nil
}

// Observe that we've modularized this code into reverse and is Equa 1 functions. We've also created a new
// class so that we can return both the head and the tail of this method. We could have also returned a two­
// element array, but that approach is less maintainable.

// Solution #2: Iterative Approach
// We want to detect linked lists where the front half of the list is the reverse of the second half. How would we
// do that? By reversing the front half of the list. A stack can accomplish this.
// We need to push the first half of the elements onto a stack. We can do this in two different ways, depending
// on whether or not we know the size of the linked list.
// If we know the size of the linked list, we can iterate through the first half of the elements in a standard for
// loop, pushing each element onto a stack. We must be careful, of course, to handle the case where the length
// of the linked list is odd.
// If we don't know the size of the linked list, we can iterate through the linked list, using the fast runner/ slow
// runner technique described in the beginning of the chapter. At each step in the loop, we push the data from
// the slow runner onto a stack. When the fast runner hits the end of the list, the slow runner will have reached
// the middle of the linked list. By this point, the stack will have all the elements from the front of the linked
// list, but in reverse order.

// Now, we simply iterate through the rest of the linked list. At each iteration, we compare the node to the top
// of the stack. If we complete the iteration without finding a difference, then the linked list is a palindrome.
func isPalindromeV2(head *Node) bool {
	fast := head
	slow := head
	stack := &stack{}
	// Push elements from first half of linked list onto stack. When fast runner
	// * (which is moving at 2x speed) reaches the end of the linked list, then we
	// * know we're at the middle
	for fast != nil && fast.next != nil {
		stack.push(slow.data)
		slow = slow.next
		fast = fast.next.next
	}
	// Has odd number of elements, so skip the middle element
	if fast != nil {
		slow = slow.next
	}
	fmt.Println("stack")
	stack.print()
	for slow != nil {
		top := stack.pop()
		// If values are different, then it's not a palindrome
		if top != slow.data {
			return false
		}
		slow = slow.next
	}
	return true
}

type stack struct {
	node  *Node
	count int
}

func (s *stack) push(data int) {
	node := &Node{data: data, next: s.node}
	s.node = node
	s.count++
}

func (s *stack) pop() int {
	if s.count == 0 {
		return -99999
	}
	data := s.node.data
	s.node = s.node.next
	s.count--
	return data
}
func (s *stack) print() {
	if s.node!=nil{
		s.node.print()
	}
}

// Solution #3: Recursive Approach
// at page 229

func main() {
	head := &Node{data: 0}
	cur := head
	for i := 1; i < 6; i++ {
		cur.next = &Node{data: i}
		cur = cur.next
	}
	for i := 4; i >= 0; i-- {
		cur.next = &Node{data: i}
		cur = cur.next
	}
	head.print()
	fmt.Println(isPalindrome(head))
	wrongHead := &Node{data: 10, next: head}
	wrongHead.print()
	fmt.Println(isPalindrome(wrongHead))

	fmt.Println("isPalindromeV2")
	head.print()
	fmt.Println(isPalindromeV2(head))
	wrongHead.print()
	fmt.Println(isPalindromeV2(wrongHead))

}
