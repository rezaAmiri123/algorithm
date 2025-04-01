package main

import "fmt"

// � Graphs
// A tree is actually a type of graph, but not all graphs are trees. Simply put, a tree is a connected graph without
// cycles.
// A graph is simply a collection of nodes with edges between (some of) them.
// •
// Graphs can be either directed (like the following graph) or undirected. While directed edges are like a
// one-way street, undirected edges are like a two-way street.
// •
// The graph might consist of multiple isolated subgraphs. If there is a path between every pair of vertices,
// it is called a "connected graph:'
// The graph can also have cycles (or not). An "acyclic graph" is one without cycles.
// Visually, you could draw a graph like this:

// In terms of programming, there are two common ways to represent a graph.

// Adjacency List
// This is the most common way to represent a graph. Every vertex (or node) stores a list of adjacent vertices.
// In an undirected graph, an edge like (a, b) would be stored twice: once in a's adjacent vertices and once
// in b's adjacent vertices.
// A simple class definition for a graph node could look essentially the same as a tree node.

type Graph struct{
	nodes []*Node
}

// The Graph class is used because, unlike in a tree, you can't necessarily reach all the nodes from a single node.
// You don't necessarily need any additional classes to represent a graph. An array (or a hash table) of lists
// (arrays, arraylists, linked lists, etc.) can store the adjacency list. The graph above could be represented as:
// 0: 1
// 1: 2
// 2: 0, 3
// 3: 2
// 4: 6
// 5: 4
// 6: 5
// This is a bit more compact, but it isn't quite as clean. We tend to use node classes unless there's a compelling
// reason not to.

// Adjacency Matrices
// An adjacency matrix is an NxN boolean matrix (where N is the number of nodes), where a true value at
// matrix[ i] [j] indicates an edge from node i to node j. (You can also use an integer matrix with Os and
// 1 s.)
// In an undirected graph, an adjacency matrix will be symmetric. In a directed graph, it will not (necessarily)
// be.

// The same graph algorithms that are used on adjacency lists (breadth-first search, etc.) can be performed
// with adjacency matrices, but they may be somewhat less efficient. In the adjacency list representation, you
// can easily iterate through the neighbors of a node. In the adjacency matrix representation, you will need to
// iterate through all the nodes to identify a node's neighbors.

// � Graph Search
// The two most common ways to search a graph are depth-first search and breadth-first search.
// In depth-first search (DFS), we start at the root (or another arbitrarily selected node) and explore each
// branch completely before moving on to the next branch. That is, we go deep first (hence the name depth­
// first search) before we go wide.
// In breadth-first search (BFS), we start at the root (or another arbitrarily selected node) and explore each
// neighbor before going on to any of their children. That is, we go wide (hence breadth-first search) before
// we go deep.
// See the below depiction of a graph and its depth-first and breadth-first search (assuming neighbors are
// iterated in numerical order).

// Breadth-first search and depth-first search tend to be used in different scenarios. DFS is often preferred if we
// want to visit every node in the graph. Both will work just fine, but depth-first search is a bit simpler.
// However, if we want to find the shortest path (or just any path) between two nodes, BFS is generally better.
// Consider representing all the friendships in the entire world in a graph and trying to find a path of friend­
// ships between Ash andVanessa.
// In depth-first search, we could take a path like Ash -> Brian -> Car let on -> Davis -> Eric
// -> Farah -> Gayle -> Harry -> Isabella -> John·-> Kari... and thenfind ourselves very
// far away. We could go through most of the world without realizing that, in fact, Vanessa is Ash's friend. We
// will still eventually find the path, but it may take a long time. It also won't find us the shortest path.
// In breadth-first search, we would stay close to Ash for as long as possible. We might iterate through many
// of Ash's friends, but we wouldn't go to his more distant connections until absolutely necessary. lfVanessa
// is Ash's friend, or his friend-of-a-friend, we'll find this out relatively quickly.

// Depth-First Search (DFS)
// In DFS, we visit a node a and then iterate through each of a's neighbors. When visiting a node b that is a
// neighbor of a, we visit all of b's neighbors before going on to a's other neighbors. That is, a exhaustively
// searches b's branch before any of its other neighbors.
// Note that pre-order and other forms of tree traversal are a form of DFS. The key difference is that when
// implementing this algorithm for a graph, we must check if the node has been visited. If we don't, we risk
// getting stuck in an infinite loop.
// The pseudocode below implements DFS.
type GraphNode struct{
	name string
	adjacent []*GraphNode
	visted bool
} 
func visitGraph(n *GraphNode){
	fmt.Println("visit node: ", n.name)
}

func dfsSearch(root *GraphNode){
	if root == nil{
		return
	}
	visitGraph(root)
	root.visted = true
	for _, n := range root.adjacent{
		if n.visted==false{
			dfsSearch(n)
		}
	}
}

// Breadth-First Search (BFS)
// BFS is a bit less intuitive, and many interviewees struggle with the implementation unless they are already
// familiar with it. The main tripping point is the (false) assumption that BFS is recursive. It's not. Instead, it
// uses a queue.
// In BFS, node a visits each of a's neighbors before visiting any of their neighbors. You can think of this as
// searching level by level out from a. An iterative solution involving a queue usually works best.
type QueueNode struct{
	next *QueueNode
	data *GraphNode
}

type Queue struct{
	head *QueueNode
	tail *QueueNode
	size int
}
func (q *Queue)enqueue(n *GraphNode){
	node := &QueueNode{data: n}
	if q.size == 0{
		q.head = node
		q.tail = node
		q.head.next = q.tail
		q.size++
		return
	}
	q.tail.next = node
	q.tail = node
	q.size++
}
func (q *Queue)dequeue()*GraphNode{
	n := q.head
	q.head = q.head.next
	q.size--
	return n.data
}
func (q *Queue)isEmpty()bool{
	if q.size == 0{
		return true
	}
	return false
}

func bfsSearch(root *GraphNode){
	queue := &Queue{}
	root.visted = true
	queue.enqueue(root) // Add to the end of queue

	for !queue.isEmpty(){
		r := queue.dequeue()// Remove from the front of the queue
		visitGraph(r)
		for _, n := range r.adjacent{
			if n.visted ==false{
				n.visted = true
				queue.enqueue(n)
			}
		}
	}
}

// If you are asked to implement BFS, the key thing to remember is the use of the queue. The rest of the algo­
// rithm flows from this fact.
// Bidirectional Search
// Bidirectional search is used to find the shortest path between a source and destination node. It operates
// by essentially running two simultaneous breadth-first searches, one from each node. When their searches
// collide, we have found a path.

// To see why this is faster, consider a graph where every node has at most k adjacent nodes and the shortest
// path from node s to nodet has length d.
// In traditional breadth-first search, we would search up to k nodes in the first "level" of the search. In the
// second level, we would search up to k nodes for each of those first k nodes, so k2 nodes total (thus far).
// We would do this d times, so that's 0( kd ) nodes.
// In bidirectional search, we have two searches that collide after approximately � levels (the midpoint
// of the path). The search from s visits approximately kd12, as does the search fromt.That's approximately
// 2 kdl2, or 0( kd/2), nodes total.
// This might seem like a minor difference, but it's not. It's huge. Recall that ( kd12 ) * ( kd12) = kd .The bidirec­
// tional search is actually faster by a factor of kd12.
// Put another way: if our system could only support searching "friend of friend" paths in breadth-first search,
// it could now likely support "friend of friend of friend of friend" paths. We can support paths that are twice
// as long.
// Additional Reading: Topological Sort (pg 632), Dijkstra's Algorithm (pg 633), AVL Trees (pg 637), Red­
// BlackTrees (pg 639).