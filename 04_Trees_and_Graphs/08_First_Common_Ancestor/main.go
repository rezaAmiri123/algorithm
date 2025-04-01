package main

// First Common Ancestor: Design an algorithm and write code to find the first common ancestor
// of two nodes in a binary tree. Avoid storing additional nodes in a data structure. NOTE: This is not
// necessarily a binary search tree.

// SOLUTION
// If this were a binary search tree, we could modify the find operation for the two nodes and see where the
// paths diverge. Unfortunately, this is not a binary search tree, so we must try other approaches.
// Let's assume we're looking for the common ancestor of nodes p and q. One question to ask here is if each
// node in our tree has a link to its parents.

// Solution #1: With Links to Parents
// If each node has a link to its parent, we could trace p and q's paths up until they intersect. This is essentially
// the same problem as question 2.7 which find the intersection of two linked lists. The "linked list" in this case
// is the path from each node up to the root. (Review this solution on page 221.)

type TreeNode struct {
	parent *TreeNode
}

func depth(node *TreeNode) int {
	depth := 0
	for node != nil {
		node = node.parent
		depth++
	}
	return depth
}
func goUpBy(node *TreeNode, delta int) *TreeNode {
	for delta > 0 && node != nil {
		node = node.parent
		delta--
	}
	return node
}

func commonAncestor(p, q *TreeNode) *TreeNode {
	delta := depth(p) - depth(q) // get difference in depths
	first := p
	second := q
	if delta < 0 {
		first, second = q, p
		delta = delta * -1
	}
	second = goUpBy(second, delta) // move deeper node up
	// Find where paths intersect.
	for first != second && first != nil && second != nil {
		first = first.parent
		second = second.parent
	}
	if first == nil || second == nil {
		return nil
	}

	return first
}
// This approach will take 0( d) time, where d is the depth of the deeper node.
// to be continued at page 268
