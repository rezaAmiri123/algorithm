package main

type TreeNode struct{
	left *TreeNode
	right *TreeNode
	value int
	visited bool
}

// � Binary Tree Traversal
// Prior to your interview, you should be comfortable implementing in-order, post-order, and pre-order
// traversal. The most common of these is in-order traversal.

// In-Order Traversal
// In-order traversal means to "visit" (often, print) the left branch, then the current node, and finally, the right
// branch
func inOrderTraversal(node *TreeNode){
	if node!= nil{
		inOrderTraversal(node.left)
		node.visited = true
		inOrderTraversal(node.right)
	}
}
// When performed on a binary search tree, it visits the nodes in ascending order (hence the name "in-order").

// Pre-Order Traversal
// Pre-order traversal visits the current node before its child nodes (hence the name "pre-order").
func preOrderTraversal(node *TreeNode){
	if node!= nil{
		node.visited = true
		preOrderTraversal(node.left)
		preOrderTraversal(node.right)
	}
}
// In a pre-order traversal, the root is always the first node visited.

// Post-Order Traversal
// Post-order traversal visits the current node after its child nodes (hence the name "post-order").
func postOrderTraversal(node *TreeNode){
	if node!= nil{
		postOrderTraversal(node.left)
		postOrderTraversal(node.right)
		node.visited = true
	}
}
// In a post-order traversal, the root is always the last node visited.

// � Binary Heaps (Min-Heaps and Max-Heaps)
// We'll just discuss min-heaps here. Max-heaps are essentially equivalent, but the elements are in descending
// order rather than ascending order.
// A min-heap is a complete binary tree (that is, totally filled other than the rightmost elements on the last
// level) where each node is smaller than its children. The root, therefore, is the minimum element in the tree.

// We have two key operations on a min-heap: insert and extract_min.
// Insert
// When we insert into a min-heap, we always start by inserting the element at the bottom. We insert at the
// rightmost spot so as to maintain the complete tree property.
// Then, we "fix"the tree by swapping the new element with its parent, until we find an appropriate spot for
// the element. We essentially bubble up the minimum element.

// This takes O( log n) time, where n is the number of nodes in the heap.
// Extract Minimum Element
// Finding the minimum element of a min-heap is easy: it's always at the top. The trickier part is how to remove
// it. (In fact, this isn't that tricky.)
// First, we remove the minimum element and swap it with the last element in the heap (the bottommost,
// rightmost element). Then, we bubble down this element, swapping it with one of its children until the min­
// heap property is restored.
// Do we swap it with the left child or the right child? That depends on their values. There's no inherent
// ordering between the left and right element, but you'll need to take the smaller one in order to maintain
// the min-heap ordering.
// This algorithm will also take 0( log n) time.

// � Tries (Prefix Trees)
// A trie (sometimes called a prefix tree) is a funny data structure. It comes up a lot in interview questions, but
// algorithm textbooks don't spend much time on this data structure.
// A trie is a variant of an n-ary tree in which characters are stored at each node. Each path down the tree may
// represent a word.
// The * nodes (sometimes called "null nodes") are often used to indicate complete words. For example, the
// fact that there is a * node under MANY indicates that MANY is a complete word. The existence of the MA path
// indicates there are words that start with MA.
// The actual implementation of these * nodes might be a special type of child (such as a
// TerminatingTrieNode, which inherits from TrieNode). Or, we could use just a boolean flag
// terminates within the "parent" node.
// A node in a trie could have anywhere from 1 through ALPHABET_SIZE + 1 children (or, 0 through
// ALPHABET_SIZE if a boolean flag is used instead of a* node).

// Very commonly, a trie is used to store the entire (English) language for quick prefix lookups. While a hash
// table can quickly look up whether a string is a valid word, it cannot tell us if a string is a prefix of any valid
// words. A trie can do this very quickly.
// How quickly? A trie can check if a string is a valid prefix in 0( K) time, where K is the length of the
// string. This is actually the same runtime as a hash table will take. Although we often refer to hash
// table lookups as being 0(1) time, this isn't entirely true. A hash table must read through all the
// characters in the input, which takes O ( K) time in the case of a word lookup.
// Many problems involving lists of valid words leverage a trie as an optimization. In situations when we search
// through the tree on related prefixes repeatedly (e.g., looking up M, then MA, then MAN, then MANY), we might
// pass around a reference to the current node in the tree. This will allow us to just check if Y is a child of MAN,
// rather than starting from the root each time.

