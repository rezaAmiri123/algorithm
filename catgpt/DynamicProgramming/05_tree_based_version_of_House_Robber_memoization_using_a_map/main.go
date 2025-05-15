package main

import "fmt"

// 🧠 Why Use a Map for Memoization?
// In the previous solution, we avoided redundant computation by returning two values
// (rob, skip) from each subtree. But in some interview scenarios,
// you may want to use a more classic top-down + memoization strategy.

// Here, we’ll define a recursive function:
// func robFrom(node *TreeNode, memo map[*TreeNode]int) int

// At each node:
// Option 1: Rob it → add node value + rob from grandchildren
// Option 2: Skip it → rob from children
// Store the max of both in the map so we don’t recompute.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	return robFrom(root, memo)
}

func robFrom(node *TreeNode, memo map[*TreeNode]int) int {
	if node == nil {
		return 0
	}

	if val, found := memo[node]; found {
		return val
	}

	// Rob this node
	moneyWithRoot := node.Val
	if node.Left != nil {
		moneyWithRoot += robFrom(node.Left.Left, memo) + robFrom(node.Left.Right, memo)
	}

	if node.Right != nil {
		moneyWithRoot += robFrom(node.Right.Left, memo) + robFrom(node.Right.Right, memo)
	}

	// Skip this node
	moneyWithoutRoot := robFrom(node.Left, memo) + robFrom(node.Right, memo)

	result := max(moneyWithRoot, moneyWithoutRoot)
	memo[node] = result
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Demo tree
func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 1}

	fmt.Println("Max money you can rob from the tree:", rob(root)) // Output: 7
}

// 🔍 How Memoization Works Here:
// We use a map[*TreeNode]int as a cache:
// If we’ve already computed the max money for a node, just return the cached value.
// This prevents recalculating subtrees, saving time.

// This pattern is widely used in:
// DP on trees
// Caching recursive results for overlapping subtrees
// Problems involving pointer-based structures (like linked lists or graphs)

// Would you like to move on to a more advanced tree DP like Binary Tree Cameras, 
// or switch to another category like Knapsack-style DP?
