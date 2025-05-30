package main

import (
	"fmt"
	"math/rand"
)

// Loop Detection: Given a circular linked list, implement an algorithm that returns the node at the
// beginning of the loop.
// DEFINITION
// Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so
// as to make a loop in the linked list.
// EXAMPLE
// Input:A - > B - > C - > D - > E - > C [the same C as earlier]
// Output:C

// SOLUTION
// This is a modification of a classic interview problem: detect if a linked list has a loop. Let's apply the Pattern
// Matching approach.

// Part 1 : Detect If Linked List Has A Loop
// An easy way to detect if a linked list has a loop is through the FastRunner / SlowRunner approach.
// FastRunner moves two steps at a time, while SlowRunner moves one step. Much like two cars racing
// around a track at different steps, they must eventually meet.

// An astute reader may wonder if FastRunner might "hop over" SlowRunner completely, without
// ever colliding. That's not possible. Suppose that FastRunner did hop over SlowRunner, such that
// SlowRunner is at spot i and FastRunner is at spot i + 1. In the previous step, SlowRunner would
// be at spot i - 1 and FastRunner would at spot ( ( i + 1) - 2), or spot i - 1. That is, they would
// have collided.

// Part 2: When Do They Collide?
// Let's assume that the linked list has a "non-looped" part of size k.
// If we apply our algorithm from part l, when will FastRunner and SlowRunner collide?
// We know that for every p steps that SlowRunner takes, FastRunner has taken 2p steps. Therefore, when
// SlowRunner enters the looped portion after k steps, FastRunner has taken 2k steps total and must be
// 2k - k steps, or k steps, into the looped portion. Since k might be much larger than the loop length, we
// should actually write this as mod ( k, LOOP_SIZE) steps, which we will denote as K.
// At each subsequent step, FastRunner and SlowRunner get either one step farther away or one step
// closer, depending on your perspective. That is, because we are in a circle, when A moves q steps away from
// B, it is also moving q steps closer to B.
// So now we know the following facts:

// 1. SlowRunner is O steps into the loop.
// 2. FastRunner is K steps into the loop.
// 3. SlowRunner is K steps behind FastRunner.
// 4. FastRunner is LOOP_SIZE - K steps behind SlowRunner.
// 5. FastRunner catches up to SlowRunner at a rate of 1 step per unit of time.
// So, when do they meet? Well, if FastRunner is LOOP_SIZE - K steps behind SlowRunner, and
// FastRunner catches up at a rate of 1 step per unit of time, then they meet after L OOP_SIZE - K steps.
// At this point, they will be K steps before the head of the loop. Let's call this point Collisions pot.

// Part 3: How Do You Find The Start of the Loop?
// We now know that CollisionSpot is K nodes before the start of the loop. Because K = mod (k, LOOP_
// SIZE) (or, in other words, k = K + M * LOOP_SIZE, for any integer M), it is also correct to say that it is
// k nodes from the loop start. For example, if node N is 2 nodes into a 5 node loop, it is also correct to say that
// it is 7, 12, or even 397 nodes into the loop.
// Therefore, both CollisionSpot and LinkedlistHead are k nodes from the start of the loop.

// Now, if we keep one pointer at CollisionSpot and move the other one to LinkedListHead, they will
// each be k nodes from LoopStart. Moving the two pointers at the same speed will cause them to collide
// again-this time after k steps, at which point they will both be at LoopStart. All we have to do is return
// this node.

// Part 4: Putting It All Together
// To summarize, we move FastPointer twice as fast as SlowPointer. When SlowPointer enters
// the loop, after k nodes, FastPointer is k nodes into the loop. This means that FastPointer and
// SlowPointer are LOOP_SIZE - k nodes away from each other.
// Next, if FastPointer moves two nodes for each node that SlowPointer moves, they move one node
// closer to each other on each turn. Therefore, they will meet after LOOP_SIZE - k turns. Both will be k
// nodes from the front of the loop.

// The head of the linked list is also k nodes from the front of the loop. So, if we keep one pointer where it is,
// and move the other pointer to the head of the linked list, then they will meet at the front of the loop.
// Our algorithm is derived directly from parts 1, 2 and 3.
// 1. Create two pointers, FastPointer and SlowPointer.
// 2. Move FastPointer at a rate of 2 steps and SlowPointer at a rate of 1 step.
// 3. When they collide, move SlowPointer to LinkedListHead. Keep FastPointer where it is.
// 4. Move SlowPointer and FastPointer at a rate of one step. Return the new collision point.
// The code below implements this algorithm.

type Node struct {
	next *Node
	data int
}

func (n *Node)print(){
	cur := n
	for i:=0;i<100&&cur!= nil;i++{
		fmt.Printf("%d-->",cur.data)
		cur=cur.next
	}
	fmt.Println()
}

func findBeginning(head *Node) *Node {
	slow := head
	fast := head

	// Find meeting point. This will be LOOP_SIZE - k steps into the linked list.
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast { //Collision
			break
		}
	}
	// Error check - no meeting point, and therefore no loop
	if fast == nil || fast.next == nil {
		return nil
	}

	// 	Move slow to Head. Keep fast at Meeting Point. Each are k steps from the
	//  Loop Start. If they move at the same pace, they must meet at Loop Start .
	slow = head
	for slow!= fast{
		slow = slow.next
		fast = fast.next
	}
	// Both now point to the start of the loop.
	return fast
}

func main(){
	head := &Node{data: 10}
	cur := head
	for i:=0;i<5;i++{
		cur.next = &Node{data: rand.Intn(10)}
		cur = cur.next
	}
	collidedPoint := cur
	for i:=0;i<5;i++{
		cur.next = &Node{data: rand.Intn(10)}
		cur = cur.next
	}
	cur.next = collidedPoint
	head.print()

	fmt.Println("collidedPoint: ",collidedPoint.data)
	result := findBeginning(head)
	fmt.Println("result: ", result.data)
}