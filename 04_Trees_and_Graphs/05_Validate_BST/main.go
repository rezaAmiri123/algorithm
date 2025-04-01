package main


// Validate BST: Implement a function to check if a binary tree is a binary search tree.
// SOLUTION
// We can implement this solution in two different ways. The first leverages the in-order traversal, and the
// second builds off the property that left <= c urrent < right.

// Solution #1: In-Order Traversal
// Our first thought might be to do an in-order traversal, copy the elements to an array, and then check to see
// if the array is sorted. This solution takes up a bit of extra memory, but it works-mostly.
// The only problem is that it can't handle duplicate values in the tree properly. For example, the algorithm
// cannot distinguish between the two trees below (one of which is invalid) since they have the same in-order
// traversal.
// However, if we assume that the tree cannot have duplicate values, then this approach works. The pseudo­
// code for this method looks something like:
type TreeNode struct{
	left,right *TreeNode
	data int
}

var index = 0
func copyBST(root *TreeNode, array []int){
	if root==nil{
		return
	}
	copyBST(root.left.left,array)
	array[index] = root.data
	index++
	copyBST(root.right, array)
}
func checkBST(root *TreeNode)bool{
	size := 5
	array := make([]int, size)
	copyBST(root, array)
	for i:=1;i<len(array);i++{
		if array[i]<=array[[i-1]]{
			return false
		}
	}
	return true
}

// Note that it is necessary to keep track of the logical "end" of the array, since it would be allocated to hold all
// the elements.
// When we examine this solution, we find that the array is not actually necessary. We never use it other than
// to compare an element to the previous element. So why not just track the last element we saw and compare
// it as we go?
// The code below implements this algorithm.

// to be continued at page 257
