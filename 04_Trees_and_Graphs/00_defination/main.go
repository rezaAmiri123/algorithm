package main

// A very simple class definition for Node is:
type Node struct{
	name string
	children []*Node
} 

// You might also have a Tree class to wrap this node. For the purposes of interview questions, we typically
// do not use a Tree class. You can if you feel it makes your code simpler or better, but it rarely does.
type Tree struct{
	root *Node
} 

// Trees vs. Binary Trees
// A binary tree is a tree in which each node has up to two children. Not all trees are binary trees

// Binary Tree vs. Binary Search Tree
// A binary search tree is a binary tree in which every node fits a specific ordering property: all left
// descendents <= n < all right descendents. This must be true for each node n.

// The definition of a binary search tree can vary slightly with respect to equality. Under some defi­
// nitions, the tree cannot have duplicate values. In others, the duplicate values will be on the right
// or can be on either side. All are valid definitions, but you should clarify this with your interviewer.

// When given a tree question, many candidates assume the interviewer means a binary search tree. Be sure
// to ask. A binary search tree imposes the condition that, for each node, its left descendents are less than or
// equal to the current node, which is less than the right descendents.

// Balanced vs. Unbalanced
// While many trees are balanced, not all are. Ask your interviewer for clarification here. Note that balancing a
// tree does not mean the left and right subtrees are exactly the same size
// One way to think about it is that a "balanced" tree really means something more like "not terribly imbal­
// anced:' It's balanced enough to ensure 0( log n) times for insert and find, but it's not necessarily as
// balanced as it could be.
// Two common types of balanced trees are red-black trees (pg 639) and AVL trees (pg 637). These are
// discussed in more detail in the Advanced Topics section.

// Complete Binary Trees
// A complete binary tree is a binary tree in which every level of the tree is fully filled, except for perhaps the
// last level. To the extent that the last level is filled, it is filled left to right.

// Full Binary Trees
// A full binary tree is a binary tree in which every node has either zero or two children. That is, no nodes have
// only one child.

// Perfect Binary Trees
// A perfect binary tree is one that is both full and complete. All leaf nodes will be at the same level, and this
// level has the maximum number of nodes.

// Note that perfect trees are rare in interviews and in real life, as a perfect tree must have exactly 2k - 1 nodes
// (where k is the number of levels). In an interview, do not assume a binary tree is perfect.
