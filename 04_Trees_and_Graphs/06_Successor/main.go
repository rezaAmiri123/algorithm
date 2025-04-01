package main

// Successor: Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a
// binary search tree. You may assume that each node has a link to its parent.

// SOLUTION
// Recall that an in-order traversal traverses the left subtree, then the current node, then the right subtree. To
// approach this problem, we need to think very, very carefully about what happens.
// Let's suppose we have a hypothetical node. We know that the order goes left subtree, then current side,
// then right subtree. So, the next node we visit should be on the right side.
// But which node on the right subtree? It should be the first node we'd visit if we were doing an in-order
// traversal of that subtree. This means that it should be the leftmost node on the right subtree. Easy enough!
// But what if the node doesn't have a right subtree? This is where it gets a bit trickier.
// If a node n doesn't have a right subtree, then we are done traversing n's subtree. We need to pick up where
// we left off with n's parent, which we'll call q.

// If n was to the left of q, then the next node we should traverse should be q (again, since left - > current
// -> right).
// If n were to the right of q, then we have fully traversed q's subtree as well. We need to traverse upwards from
// q until we find a node x that we have not fully traversed. How do we know that we have not fully traversed
// a node x? We know we have hit this case when we move from a left node to its parent. The left node is fully
// traversed, but its parent is not

// The pseudocode looks like this:
// Node inorderSucc(Node n) {
// 	if (n has a right subtree) {
// 		return leftmost child of right subtree
// 	} else {
// 		while (n is a right child of n.parent) {
// 			n = n.parent; // Go up
// 		}
// 		return n.parent; // Parent has not been traversed
// 	}
// }

// But wait-what if we traverse all the way up the tree before finding a left child?This will happen only when
// we hit the very end of the in-order traversal. That is, if we're already on the far right of the tree, then there is
// no in-order successor. We should return null.
// The code below implements this algorithm (and properly handles the null case)
type TreeNode struct {
	left, right, parent *TreeNode
	value               int
}

func inorderSucc(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}
	// Found right children -> return leftmost node of right subtree.
	if n.right != nil {
		return leftMostChild(n.right)
	} else {
		q := n
		x := q.parent
		// Go up until we're on left instead of right
		for x != nil && x.left != q {
			q = x
			x = x.parent
		}
		return x
	}
}

func leftMostChild(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

// This is not the most algorithmically complex problem in the world, but it can be tricky to code perfectly. In
// a problem like this, it's useful to sketch out pseudocode to carefully outline the different cases.
