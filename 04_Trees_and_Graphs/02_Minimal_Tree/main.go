package main

import "fmt"

// Minimal Tree: Given a sorted (increasing order) array with unique integer elements, write an
// algorithm to create a binary search tree with minimal height

// SOLUTION
// To create a tree of minimal height, we need to match the number of nodes in the left subtree to the number
// of nodes in the right subtree as much as possible. This means that we want the root to be the middle of the
// array, since this would mean that half the elements would be less than the root and half would be greater
// than it.
// We proceed with constructing our tree in a similar fashion. The middle of each subsection of the array
// becomes the root of the node. The left half of the array will become our left subtree, and the right half of
// the array will become the right subtree.
// One way to implement this is to use a simple root. insertNode(int v) method which inserts the
// value v through a recursive process that starts with the root node. This will indeed construct a tree with
// minimal height but it will not do so very efficiently. Each insertion will require traversing the tree, giving a
// total cost ofO( N log N) to the tree.
// Alternatively, we can cut out the extra traversals by recursively using the createMinimalBST method.
// This method is passed just a subsection of the array and returns the root of a minimal tree for that array.
// The algorithm is as follows:
// 1. Insert into the tree the middle element of the array.
// 2. Insert (into the left subtree) the left subarray elements.
// 3. Insert (into the right subtree) the right subarray elements.
// 4. Recurse.
// The code below implements this algorithm.

type TreeNode struct{
	left *TreeNode
	right *TreeNode
	value int
}

func(n *TreeNode)print(){
	if n.left!= nil{
		n.left.print()
	}
	fmt.Printf("%d-->", n.value)
	if n.right!= nil{
		n.right.print()
	}
}

func CreateMinimalBST(array []int)*TreeNode{
	return createMinimalBST(array,0,len(array)-1)
}
func createMinimalBST(arr []int,start,end int)*TreeNode{
	if end<start{
		return nil
	}
	mid :=(start+end)/2
	n := &TreeNode{value: arr[mid]}
	n.left=createMinimalBST(arr,start,mid-1)
	n.right=createMinimalBST(arr,mid+1,end)
	return n
}

// Although this code does not seem especially complex, it can be very easy to make little off-by-one errors.
// Be sure to test these parts of the code very thoroughly.

func main(){
	array := []int{}
	for i := 0;i<10;i++{
		array = append(array, i)
	}
	bst := CreateMinimalBST(array)
	bst.print()
	fmt.Println()
}





