package main

import (
	"fmt"
	"math"
)

// Great pick! 🌳 An AVL Tree is a self-balancing Binary Search Tree (BST) —
// often asked in system design-heavy interviews or advanced algorithm rounds.

// ✅ What’s an AVL Tree?
// A BST that maintains balance by making sure the difference in height between left
// and right subtrees of any node is at most 1.

// When an imbalance happens due to insertion or deletion, rotations are performed:
// Left Rotation
// Right Rotation
// Left-Right Rotation
// Right-Left Rotation

// 🔧 AVL Tree Insert Example in Go
// Here’s a minimal AVL Tree in Go that supports insertions and balances itself:

// Node represents a node in the AVL tree
type Node struct{
	Value int
	Left *Node
	Right *Node
	Height int
}

// Get height of node
func height(n *Node)int{
	if n==nil{
		return 0
	}
	return n.Height
}

// Get balance factor
func getBalance(n *Node)int{
	if n == nil{
		return 0
	}
	return height(n.Left)-height(n.Right)
}

// Right rotate
func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right
	//     y
	//   x
	//     T2
	x.Right = y
	y.Left = T2
	// x
	//   y
	// T2
	y.Height = int(math.Max(float64(height(y.Left)), float64(height(y.Right)))) + 1// T2 - ?
	x.Height = int(math.Max(float64(height(x.Left)), float64(height(x.Right)))) + 1// ? - y

	return x
}

// Left rotate
func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left
	//  x
	//    y
	// T2
	y.Left = x
	x.Right = T2
	//   y
	// x 
	//    T2
	x.Height = int(math.Max(float64(height(x.Left)), float64(height(x.Right)))) + 1 // ? - T2
	y.Height = int(math.Max(float64(height(y.Left)), float64(height(y.Right)))) + 1 // x - ?

	return y
}

// Insert into AVL tree
func insert(root *Node, key int)*Node{
	if root == nil{
		return &Node{Value: key, Height: 1}
	}

	if key<root.Value{
		root.Left = insert(root.Left, key)
	}else if key>root.Value{
		root.Right = insert(root.Right,key)
	}else{
		// Duplicate keys not allowed
		return root
	}

	// Update height
	root.Height = 1+int(math.Max(float64(height(root.Left)),float64(height(root.Right))))

	// Balance
	balance := getBalance(root)

	// Left Left
	if balance>1&&key<root.Left.Value{
		return rightRotate(root)
	}

	// Right Right
	if balance < -1&&key>root.Right.Value{
		return leftRotate(root)
	}

	// Left Right
	if balance > -1&&key>root.Left.Value{
		root.Left = leftRotate(root.Left)
		return rightRotate(root)
	}

	// Right Left
	if balance < -1 && key < root.Right.Value {
		root.Right = rightRotate(root.Right)
		return leftRotate(root)
	}

	return root
}

// InOrder traversal (sorted)
func inOrder(n *Node){
	if n!= nil{
		inOrder(n.Left)
		fmt.Printf("%d ",n.Value)
		inOrder(n.Right)
	}
}

func main(){
	var root *Node
	keys := []int{10, 20, 30, 40, 50, 25}

	for _, key := range keys{
		root = insert(root, key)
	}

	fmt.Print("InOrder traversal of the AVL Tree: ")
	inOrder(root) // Should be sorted
	fmt.Println() // InOrder traversal of the AVL Tree: 10 20 25 30 40 50
}

// You can extend this with:
// Delete operation
// Search
// Pre/Post-Order Traversal
// Understanding when and why each rotation is triggered is key.
// Want to try building delete next? Or maybe use AVL tree for interval or range queries?