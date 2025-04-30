package main

import "fmt"

// ✅ Problem: Detect Cycle in an Undirected Graph
// Description:
// Given n nodes and a list of edges representing an undirected graph,
// determine if the graph contains any cycle.

// 🧠 Idea:
// For an undirected graph:
// Use DFS to traverse the graph.
// Keep track of the parent node in recursion.
// If a visited node is reached and it's not the parent, there is a cycle.

func hasCycle(n int, edges [][]int) bool {
	graph := make(map[int][]int)
	visited := make([]bool, n)

	// Build the adjacency list
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var dfs func(int, int) bool
	dfs = func(node, parent int) bool {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				if dfs(neighbor, node) {
					return true
				}
			} else if neighbor != parent {
				// Found a back edge (cycle)
				return true
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			if dfs(i, -1) {
				return true
			}
		}
	}
	return false
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}}

	fmt.Println("Does the graph have a cycle?", hasCycle(n, edges)) // Output: true
	// With cycle
	edges = [][]int{{0, 1}, {1, 2}, {2, 0}}                         // true
	fmt.Println("Does the graph have a cycle?", hasCycle(n, edges)) // Output: true
	// Without cycle
	edges = [][]int{{0, 1}, {1, 2}, {3, 4}}                         // false
	fmt.Println("Does the graph have a cycle?", hasCycle(n, edges)) // Output: false
}

// If you're comfortable with this, we can move to:

// Cycle detection in directed graphs (a bit trickier!)

// Topological sort

// Or anything else you're curious about.