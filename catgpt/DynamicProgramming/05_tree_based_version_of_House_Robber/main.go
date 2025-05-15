package main

import "fmt"

// Excellent choice! The tree-based version of House Robber is known as House Robber III,
// and it's a popular medium-to-hard dynamic programming problem
// that combines DP + Trees + Memoization.

// ✅ Problem: House Robber III (LeetCode 337)
// Problem Statement:
// The thief has found himself a new place for his thievery.
// There is only one entrance to this area, called the root of the binary tree.
// Each node in the tree represents a house with a certain amount of money.
// The thief cannot rob two directly-linked houses (parent and child).
// Return the maximum amount of money the thief can rob without alerting the police.

// 🌳 Key Idea:
// At each node, you have two choices:
// Rob this node: then you cannot rob its immediate children.
// Skip this node: then you are free to rob its children.
// We define a function that returns two values:
// robThis: max money if we rob this node
// skipThis: max money if we do NOT rob this node

// 💡 Recurrence:
// For a node n, and its left/right children:

// robThis = n.Val + left.skip + right.skip
// skipThis = max(left.rob, left.skip) + max(right.rob, right.skip)

// TreeNode defines a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Pair holds the values: (robThis, skipThis)
type Pair struct {
	rob  int
	skip int
}

func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res.rob, res.skip)
}

func dfs(node *TreeNode) Pair {
	if node == nil {
		return Pair{0, 0}
	}

	left := dfs(node.Left)
	right := dfs(node.Right)

	// Rob this node: can't rob children
	robThis := node.Val + left.skip + right.skip

	// Skip this node: take max of rob/skip of children
	skipThis := max(left.rob, left.skip) + max(right.rob, right.skip)

	return Pair{robThis, skipThis}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper to build test tree
func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 1}

	fmt.Println("Max money you can rob from the tree:", rob(root)) // Output: 7
}

// 📚 Explanation:
// Tree structure:

//       3
//      / \
//     2   3
//      \   \
//       3   1
// Rob 3 (root) + 3 (left.right) + 1 (right.right) = 7
// If we rob root, we can't rob its direct children (2 and 3), but we can rob grandchildren.
// This problem showcases:
// Post-order traversal (DFS)
// DP on trees
// Elegant state management with tuple return

// Would you like to try implementing memoization on the tree using a map for practice, 
// or go for another tree-DP problem like Binary Tree Camera Placement 
// or Diameter of Binary Tree with memoization?