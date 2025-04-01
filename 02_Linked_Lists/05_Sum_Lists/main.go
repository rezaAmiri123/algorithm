package main

import (
	"fmt"
	"math/rand"
)

// Sum Lists: You have two numbers represented by a linked list, where each node contains a single
// digit. The digits are stored in reverse order,such that the 1's digit is at the head of the list. Write a
// function that adds the two numbers and returns the sum as a linked list.
// EXAMPLE
// Input: (7-> 1 -> 6) + (5 -> 9 -> 2).That is,617 + 295.
// Output: 2 -> 1 -> 9. That is,912.
// FOLLOW UP
// Suppose the digits are stored in forward order. Repeat the above problem.
// Input: (6 -> 1 -> 7) + (2 -> 9 -> 5).That is,617 + 295.
// Output: 9 -> 1 -> 2.That is, 912.

// SOLUTION
// It's useful to remember in this problem how exactly addition works. Imagine the problem:
// 6 1 7
// + 2 9 5
// First, we add 7 and 5 to get 12. The digit 2 becomes the last digit of the number, and 1 gets carried over to
// the next step. Second, we add 1, 1, and 9 to get 11. The 1 becomes the second digit,and the other 1 gets
// carried over the final step. Third and finally, we add 1,6 and 2 to get 9. So,our value becomes 912

// We can mimic this process recursively by adding node by node,carrying over any "excess" data to the next
// node. Let's walk through this for the below linked list:
//   7 -> 1 -> 6
// + 5 -> 9 -> 2

// We do the following:
// 1. We add 7 and 5 first,getting a result of 12. 2 becomes the first node in our linked list,and we "carry" the
// 1 to the next sum.
// List: 2 ->?
// 2. We then add 1 and 9, as well as the "carry;' getting a result of 11. 1 becomes the second element of our
// linked list, and we carry the 1 to the next sum.
// List: 2 -> 1 ->?
// 3. Finally, we add 6, 2 and our"carrY:'to get 9.This becomes the final element of our linked list.
// List: 2 -> 1 -> 9.
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

func addLists(l1, l2 *Node, carry int) *Node {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	result := &Node{}
	value := carry
	if l1 != nil {
		value += l1.data
	}
	if l2 != nil {
		value += l2.data
	}
	result.data = value % 10 //Second digit of number

	// Recurse
	if l1 != nil || l2 != nil {
		more := addLists(l1.next, l2.next, value/10)
		result.next = more
	}
	return result
}

// In implementing this code, we must be careful to handle the condition when one linked list is shorter than
// another. We don't want to get a null pointer exception.

// Follow Up
// Part B is conceptually the same (recurse, carry the excess), but has some additional complications when it
// comes to implementation:
// 1. One list may be shorter than the other, and we cannot handle this "on the flY:' For example, suppose we
// were adding (1 -> 2 -> 3-> 4) and (5-> 6-> 7). We need to know that the 5 should be"matched"with the
// 2, not the 1. We can accomplish this by comparing the lengths of the lists in the beginning and padding
// the shorter list with zeros.
// 2. In the first part, successive results were added to the tail (i.e., passed forward). This meant that the recur­
// sive call would be passed the carry, and would return the result (which is then appended to the tail). In
// this case, however, results are added to the head (i.e., passed backward). The recursive call must return
// the result, as before, as well as the carry. This is not terribly challenging to implement, but it is more
// cumbersome. We can solve this issue by creating a wrapper class called Partial Sum.
// The code below implements this algorithm.

type PartialSum struct {
	sum   *Node
	carry int
}

func addListsImproved(l1, l2 *Node) *Node {
	len1 := l1.length()
	len2 := l2.length()

	// Pad the shorter list with zeros - see note (1)
	if len1 < len2 {
		l1 = padList(l1, len2-len1)
	} else {
		l2 = padList(l2, len1-len2)
	}

	// Add lists
	sum := addListsHelper(l1, l2)

	/* If there was a carry value left over, insert this at the front of the list.
	* Otherwise, just return the linked list. */
	if sum.carry == 0{
		return sum.sum
	}else{
		result
	}
}

func addListsHelper(l1, l2 *Node) *PartialSum {
	if l1 == nil && l2 == nil {
		return &PartialSum{}
	}
	// Add smaller digits recursively
	sum := addListsHelper(l1.next, l2.next)

	// Add carry to current data
	val := sum.carry + l1.data + l2.data

	// Insert sum of current digits
	if sum.sum != nil {
		temp := &Node{data: val % 10, next: sum.sum}
		sum.sum = temp
		sum.carry = val / 10
	}
	return sum

}

// Pad the list with zeros
func padList(head *Node, padding int) *Node {
	for i := 0; i < padding; i++ {
		n := &Node{data: 0, next: head}
		head = n
	}
	return head
}

func (n *Node) length() int {
	length := 0
	for n != nil {
		length++
		n = n.next
	}
	return length
}

func main() {
	l1 := &Node{data: 1}
	l2 := &Node{data: 1}

	cur1 := l1
	cur2 := l2
	for i := 0; i < 10; i++ {
		cur1.next = &Node{data: rand.Intn(10)}
		cur1 = cur1.next
		cur2.next = &Node{data: rand.Intn(10)}
		cur2 = cur2.next
	}
	l1.print()
	l2.print()
	result := addLists(l1, l2, 0)
	result.print()

	// improved
	l1 = &Node{data: 1}
	l2 = &Node{data: 1}

	cur1 = l1
	cur2 = l2
	for i := 0; i < 10; i++ {
		cur1.next = &Node{data: rand.Intn(10)}
		cur1 = cur1.next
		cur2.next = &Node{data: rand.Intn(10)}
		cur2 = cur2.next
	}
	cur2.next = &Node{data: rand.Intn(10)}
	cur2 = cur2.next
	l1.print()
	l2.print()
	result = addListsImproved(l1, l2)
	result.print()
}
