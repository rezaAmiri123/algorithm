package main

// Paths with Sum: You are given a binary tree in which each node contains an integer value (which
// 	might be positive or negative). Design an algorithm to count the number of paths that sum to a
// 	given value. The path does not need to start or end at the root or a leaf, but it must go downwards
// 	(traveling only from parent nodes to child nodes).

// SOLUTION
// Let's pick a potential sum-say, 8-and then draw a binary tree based on this. This tree intentionally has a
// . number of paths with this sum.

// One option is the brute force approach.

// Solution #1: Brute Force
// In the brute force approach, we just look at all possible paths. To do this, we traverse to each node. At each
// node, we recursively try all paths downwards, tracking the sum as we go. As soon as we hit our target sum,
// we increment the total.
type TreeNode struct {
	left, right *TreeNode
	value       int
}

func countPathsWithSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// Count paths with sum starting from the root.
	pathsFromRoot := countPathsWithSumFromNode(root, targetSum, 0)

	// Try the nodes on the left and right.
	pathOnLeft := countPathsWithSum(root.left, targetSum)
	pathOnRight := countPathsWithSum(root.right, targetSum)

	return pathsFromRoot + pathOnLeft + pathOnRight
}

// Returns the number of paths with this sum starting from this node.
func countPathsWithSumFromNode(node *TreeNode, targetSum int, currentSum int) int {
	if node != nil {
		return 0
	}

	currentSum += node.value
	totalPaths := 0
	if currentSum == targetSum { // Found a path from the root
		totalPaths++
	}
	totalPaths += countPathsWithSumFromNode(node.left, targetSum, currentSum)
	totalPaths += countPathsWithSumFromNode(node.right, targetSum, currentSum)
	return totalPaths
}

// What is the time complexity of this algorithm?
// Consider that node at depth d will be "touched" (via countPathsWithSumFromNode) by d nodes above it.
// In a balanced binary tree, d will be no more than approximately log N. Therefore, we know that with N
// nodes in the tree, countPathsWithSumFromNode will be called O(N log N) times. The runtime is O(N log N).
// We can also approach this from the other direction. At the root node, we traverse to all N - 1 nodes beneath
// it (via countPathsWithSumFromNode). At the second level (where there are two nodes), we traverse to N - 3
// nodes. At the third level (where there are four nodes, plus three above those), we traverse to N - 7 nodes.
// Following this pattern, the total work is roughly:
// (N - 1) + (N - 3) + (N - 7) + (N - 15) + (N - 31) + ••• + (N - N)
// To simplify this, notice that the left side of each term is always N and the right side is one less than a power
// of two. The number of terms is the depth of the tree, which is O(log N). For the right side, we can ignore
// the fact that it's one less than a power of two. Therefore, we really have this:
// O(N * [number of terms] - [sum of powers of two from 1 through NJ)
// O(N log N - N)
// O(N log N)
// If the value of the sum of powers of two from 1 through N isn't obvious to you, think about what the powers
// of two look like in binary:
// 0001
// + 0010
// + 0100
// + 1000
// = 1111
// Therefore, the runtime is O(N log N) in a balanced tree.
// In an unbalanced tree, the runtime could be much worse. Consider a tree that is just a straight line down. At
// the root, we traverse to N - 1 nodes. At the next level (with just a single node), we traverse to N - 2 nodes.
// At the third level, we traverse to N - 3 nodes, and so on. This leads us to the sum of numbers between 1
// and N, which is 0( N2 ).

// to be continued at page 284
