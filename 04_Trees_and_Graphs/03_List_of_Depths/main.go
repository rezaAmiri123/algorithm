package main

import "fmt"

// List of Depths: Given a binary tree, design an algorithm which creates a linked list of all the nodes
// at each depth (e.g., if you have a tree with depth D, you'll have D linked lists).

// SOLUTION
// Though we might think at first glance that this problem requires a level-by-level traversal, this isn't actually
// necessary. We can traverse the graph any way that we'd like, provided we know which level we're on as we
// do so.
// We can implement a simple modification of the pre-order traversal algorithm, where we pass in level +
// 1 to the next recursive call. The code below provides an implementation using depth-first search.

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
	next  *TreeNode
}

type LinkedList struct {
	node  *TreeNode
	count int
}

func (ll *LinkedList) add(node *TreeNode) {
	ll.count++
	if ll.node == nil {
		ll.node = node
		return
	}
	node.next = ll.node
	ll.node = node
}

func createLevelLinedList(root *TreeNode, lists []*LinkedList, level int) {
	if root == nil {
		return // base case
	}

	var list *LinkedList
	if len(lists) == level { // Level not contained in list
		list = &LinkedList{}
		// I* Levels are always traversed in order. So, if this is the first time we've
		// * visited level i, we must have seen levels 0 through i - 1. We can
		// * therefore safely add the level at the end. *I
		lists = append(lists, list)
	} else {
		list = lists[level]
	}
	list.add(root)
	createLevelLinedList(root.left, lists, level+1)
	createLevelLinedList(root.right, lists, level+1)
}

func main() {
	nodes := []*TreeNode{}
	for i := 0; i < 10; i++ {
		nodes = append(nodes, &TreeNode{data: i})
	}
	nodes[5].left = nodes[3]
	nodes[5].right = nodes[7]

	nodes[3].left = nodes[2]
	nodes[4].right = nodes[4]

	lists := []*LinkedList{}
	createLevelLinedList(nodes[5],lists,0)
	fmt.Println(lists)
	// 253
}
