package main

// Rank from Stream: Imagine you are reading in a stream of integers. Periodically, you wish
// to be able to look up the rank of a number x (the number of values less than or equal to x).
// Implement the data structures and algorithms to support these operations. That is, implement
// the method track(int x), which is called when each number is generated, and the method
// getRankOfNumber(int x), which returns the number of values less than or equal to x (not
// including x itself).
// EXAMPLE
// Stream(in order of appearance):5, 1, 4, 4, 5, 9, 7, 13, 3
// getRankOfNumber(l)=0
// getRankOfNumber(3)=1
// getRankOfNumber(4)=3

// SOLUTION
// A relatively easy way to implement this would be to have an array that holds all the elements in sorted
// order. When a new element comes in, we would need to shift the other elements to make room. Imple­
// menting getRankOfNumber would be quite efficient, though. We would simply perform a binary search
// for n, and return the index.
// However, this is very inefficient for inserting elements (that is, the track(int x) function). We need a
// data structure which is good at keeping relative ordering, as well as updating when we insert new elements.
// A binary search tree can do just that.
// Instead of inserting elements into an array, we insert elements into a binary search tree. The method
// track(int x) will run in O(log n) time, where n is the size of the tree (provided, of course, that the
// tree is balanced).

// To find the rank of a number, we could do an in-order traversal, keeping a counter as we traverse. The goal
// is that, by the time we find x, counter will equal the number of elements less than x.
// As long as we're moving left during searching for x, the counter won't change. Why? Because all the values
// we're skipping on the right side are greater than x. After all, the very smallest element(with rank of 1) is the
// leftmost node.
// When we move to the right though, we skip over a bunch of elements on the left. All of these elements are
// less than x, so we'll need to increment counter by the number of elements in the left subtree.
// Rather than counting the size of the left subtree(which would be inefficient), we can track this information
// as we add new elements to the tree.
// Let's walk through an example on the following tree. In the below example, the value in parentheses indi­
// cates the number of nodes in the left subtree (or, in other words, the rank of the node relative to its subtree).

// Recursively, the algorithm is the following:
// int getRank(Node node, int x) {
// 	if x is node.data, return node.leftSize()
// 	if x is on left of node, return getRank(node.left, x)
// 	if x is on right of node, return node.leftSize() + 1 + getRank(node.right, x)
// }

// The full code for this is below.
type RankNode struct{
	leftSize ,data int
	left,right *RankNode
} 
func(n *RankNode)insert (d int){
	if d<=n.data{
		if n.left!= nil{
			n.left.insert(d)
		}else{
			n.left= &RankNode{data: d}
		}
		n.leftSize++
	}else{
		if n.right!= nil{
			n.right.insert(d)
		}else{
			n.right = &RankNode{data: d}
		}
	}
}

var root *RankNode

func track(number int){
	if root!= nil{
		root = &RankNode{data: number}
	}else{
		root.insert(number)
	}
}

func (n*RankNode)getRank(d int)int{
	if d==n.data{
		return n.leftSize
	}else if d<n.data{
		if n.left == nil{
			return -1
		}else {
			return n.left.getRank(d)
		}
	}else{
		rightRank := -1
		if n.right!= nil{
			rightRank = n.right.getRank(d)
		}
		if rightRank==-1{
			return -1
		}else{
			return n.leftSize+1+rightRank
		}
	}
}

func getRankOfNumber(number int)int{
	return root.getRank(number)
}

// The track method and the getRankOfNumber method will both operate in O(log N) on a balanced
// tree and O(N) on an unbalanced tree.
// Note how we've handled the case in which d is not found in the tree. We check for the -1 return value, and,
// when we find it, return -1 up the tree. It is important that you handle cases like this.
// 425