package main

import (
	"math/rand"

	"github.com/docker/docker/testutil/daemon"
)

// Random Node: You are implementing a binary search tree class from scratch, which, in addition
// to insert, find, and delete, has a method getRandomNode() which returns a random node
// from the tree. All nodes should be equally likely to be chosen. Design and implement an algorithm
// for getRandomNode, and explain how you would implement the rest of the methods.

// SOLUTION
// Let's draw an example.

// We're going to explore many solutions until we get to an optimal one that works.
// One thing we should realize here is that the question was phrased in a very interesting way. The interviewer
// did not simply say, "Design an algorithm to return a random node from a binary tree:'We were told that this
// is a class that we're building from scratch. There is a reason the question was phrased that way. We probably
// need access to some part of the internals of the data structure.

// Option #1 [Slow & Working]
// One solution is to copy all the nodes to an array and return a random element in the array. This solution will
// take O(N) time and O(N) space, where N is the number of nodes in the tree.
// We can guess our interviewer is probably looking for something more optimal, since this is a little too
// straightforward (and should make us wonder why the interviewer gave us a binary tree, since we don't
// need that information).
// We should keep in mind as we develop this solution that we probably need to know something about the
// internals of the tree. Otherwise, the question probably wouldn't specify that we're developing the tree class
// from scratch.

// Option #2 [Slow & Working)
// Returning to our original solution of copying the nodes to an array, we can explore a solution where we
// maintain an array at all times that lists all the nodes in the tree. The problem is that we'll need to remove
// nodes from this array as we delete them from the tree, and that will take O(N) time.

// Option #3 [Slow & Working]
// We could label all the nodes with an index from 1 to N and label them in binary search tree order (that
// is, according to its inorder traversal). Then, when we call getRandomNode, we generate a random index
// between 1 and N. If we apply the label correctly, we can use a binary search tree search to find this index.
// However, this leads to a similar issue as earlier solutions. When we insert a node or a delete a node, all of the
// indices might need to be updated. This can take O(N) time.

// Option #4 [Fast & Not Working]
// What if we knew the depth of the tree? (Since we're building our own class, we can ensure that we know
// this. It's an easy enough piece of data to track.)
// We could pick a random depth, and then traverse left/right randomly until we go to that depth. This
// wouldn't actually ensure that all nodes are equally likely to be chosen though.
// First, the tree doesn't necessarily have an equal number of nodes at each level. This means that nodes on
// levels with fewer nodes might be more likely to be chosen than nodes on a level with more nodes.
// Second, the random path we take might end up terminating before we get to the desired level. Then what?
// We could just return the last node we find, but that would mean unequal probabilities at each node.

// Option #5 [Fast & Not Working]
// We could try just a simple approach: traverse randomly down the tree. At each node:
// X odds, we return the current node.
// • With X odds, we traverse left.
// With X odds, we traverse right.
// With
// This solution, like some of the others, does not distribute the probabilities evenly across the nodes. The root
// has a
// probability of being selected-the same as all the nodes in the left put together.

// Option #6 [Fast & Working]
// Rather than just continuing to brainstorm new solutions, let's see if we can fix some of the issues in the
// previous solutions. To do so, we must diagnose-deeply-the root problem in a solution

// Let's look at Option #5. It fails because the probabilities aren't evenly distributed across the options. Can we
// fix that while keeping the basic algorithm the same?
type TreeNode struct{
	data, size int
	left, right *TreeNode
}

func (tn *TreeNode)getRandomNode()*TreeNode{
	var leftSize int
	if tn.left!= nil{
		leftSize = tn.left.size
	}
	index := rand.Intn(leftSize)
	if index<leftSize{
		return tn.left.getRandomNode()
	}else if index==leftSize{
		return tn
	}else{
		return tn.right.getRandomNode()
	}
}

func (tn *TreeNode)insertInOrder(d int){
	if d<=tn.data{
		if tn.left==nil{
			tn.left=&TreeNode{data: d,size: 1}
		} else{
			tn.right.insertInOrder(d)
		}
	}
	tn.size++
}

func (tn *TreeNode)find(d int)*TreeNode{
	if d==tn.data{
		return tn
	}else if d<=tn.data{
		if tn.left !=nil{
			return tn.left.find(d)
		}
	} else if d>tn.data{
		if tn.right!=nil{
			return tn.right.find(d)
		}
	}
	return nil
}

// balanced tree, this algorithm will be O(log N), where N is the number of nodes.

// to be continued at page 282
