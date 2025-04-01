package main

// Check Subtree:Tl and T2 are two very large binary trees, withTl much bigger thanT2. Create an
// algorithm to determine ifT2 is a subtree ofTl.
// A tree T2 is a subtree of Tl if there exists a node n in Tl such that the subtree of n is identical to T2.
// That is, if you cut off the tree at node n, the two trees would be identical.

// SOLUTION
// In problems like this, it's useful to attempt to solve the problem assuming that there is just a small amount
// of data. This will give us a basic idea of an approach that might work.
// The Simple Approach
// In this smaller, simpler problem, we could consider comparing string representations of traversals of each
// tree. lfT2 is a subtree ofTl, thenT2's traversal should be a substring ofTl. ls the reverse true? If so, should
// we use an in-order traversal or a pre-order traversal?
// An in-order traversal will definitely not work. After all, consider a scenario in which we were using binary
// search trees. A binary search tree's in-order traversal always prints out the values in sorted order.Therefore,
// two binary search trees with the same values will always have the same in-order traversals, even if their
// structure is different.
// What about a pre-order traversal?This is a bit more promising. At least in this case we know certain things,
// like the first element in the pre-order traversal is the root node.The left and right elements will follow.
// Unfortunately, trees with different structures could still have the same pre-order traversal.

// There's a simple fix though. We can store NULL nodes in the pre-order traversal string as a special character,
// like an 'X'. (We'll assume that the binary trees contain only integers.)The left tree would have the traversal
// { 3, 4, X} and the right tree will have the traversal { 3, X, 4}.
// Observe that, as long as we represent the NULL nodes, the pre-order traversal of a tree is unique.That is, if
// two trees have the same pre-order traversal, then we know they are identical trees in values and structure.

// To see this, consider reconstructing a tree from its pre-order traversal (with NULL nodes indicated). For
// example: 1, 2, 4, X, X, X, 3, X, X.
// The root is 1, and its left node, 2, follows it. 2.left must be 4. 4 must have two NULL nodes (since it is followed
// by two Xs). 4 is complete, so we move back up to its parent, 2. 2.right is another X (NULL). 1's left subtree
// is now complete, so we move to 1's right child. We place a 3 with two NULL children there.The tree is now
// complete

// This whole process was deterministic, as it will be on any other tree. A pre-order traversal always starts at
// the root and, from there, the path we take is entirely defined by the traversal. Therefore, two trees are iden­
// tical if they have the same pre-order traversal.
// Now consider the subtree problem. lfT2's pre-order traversal is a substring ofTl's pre-order traversal, then
// T2's root element must be found inTl. If we do a pre-order traversal from this element in Tl, we will follow
// an identical path toT2's traversal. Therefore,T2 is a subtree of Tl.
// Implementing this is quite straightforward. We just need to construct and compare the pre-order traversaIs

// continued at page 277
