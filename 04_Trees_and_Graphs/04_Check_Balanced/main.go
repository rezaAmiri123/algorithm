package main

import "math"

// Check Balanced: Implement a function to check if a binary tree is balanced. For the purposes of
// this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any
// node never differ by more than one.

// SOLUTION
// In this question, we've been fortunate enough to be told exactly what balanced means: that for each node,
// the two subtrees differ in height by no more than one. We can implement a solution based on this defini­
// tion. We can simply recurse through the entire tree, and for each node, compute the heights of each subtree.
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return -1 // base case
	}
	answer := getHeight(root.left)
	right := getHeight(root.right)
	if right > answer {
		answer = right
	}
	return answer + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true // base case
	}
	heightDiff := getHeight(root.left) - getHeight(root.right)
	abs := heightDiff
	if abs < 0 {
		abs = abs * -1
	}
	if heightDiff > 1 {
		return false
	} else { //Recurse
		return isBalanced(root.left) && isBalanced(root.right)
	}
}

// Although this works. it's not very efficient. On each node. we recurse through its entire subtree. This means
// that getHeight is called repeatedly on the same nodes. The algorithm isO(N log N) since each node is
// "touched" once per node above it.
// We need to cut out some of the calls to getHeight.
// If we inspect this method, we may notice that getHeight could actually check if the tree is balanced at
// the same time as it's checking heights. What do we do when we discover that the subtree isn't balanced?
// Just return an error code.
// This improved algorithm works by checking the height of each subtree as we recurse down from the root.
// On each node, we recursively get the heights of the left and right subtrees through the checkHeight
// method. If the subtree is balanced, then checkHeight will return the actual height of the subtree. If the
// subtree is not balanced, then checkHeight will return an error code. We will immediately break and
// return an error code from the current call.

// What do we use for an error code? The height of a null tree is generally defined to be -1, so that's
// not a great idea for an error code. Instead, we' ll use Integer. MIN_VALUE
func checkHeightImproved(root *TreeNode) int {
	if root == nil {
		return -1
	}
	leftHight := checkHeightImproved(root.left)
	if leftHight == math.MinInt {
		return math.MinInt // Pass error up
	}
	rightHight := checkHeightImproved(root.right)
	if rightHight == math.MinInt {
		return math.MinInt // Pass error up
	}
	absHeightDiff := leftHight - rightHight
	if absHeightDiff < 0 {
		absHeightDiff = -1 * absHeightDiff
	}
	if absHeightDiff > 1 {
		return math.MinInt
	} else {
		max := leftHight
		if rightHight > max {
			max = rightHight
		}
		return max + 1
	}
}

func isBalancedImproved(root *TreeNode)bool{
	return checkHeightImproved(root) != math.MinInt
}
// This code runs in O(N) time andO(H) space, where H is the height of the tree.
