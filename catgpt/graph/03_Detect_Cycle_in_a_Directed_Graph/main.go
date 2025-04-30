package main

import "fmt"

// ✅ Problem: Detect Cycle in a Directed Graph
// Description:
// Given n nodes and a list of directed edges, determine if there's a cycle in the graph.

// 🧠 Idea:
// To detect a cycle in a directed graph, you can use DFS + recursion stack:
// Keep a visited array for visited nodes.
// Keep a recStack (recursion stack) to track nodes in the current DFS path.
// If a node is already in the recursion stack, there's a cycle.

func hasDirectedCycle(n int, edges [][]int) bool {
	graph := make(map[int][]int)
	visited := make([]bool, n)
	recStack := make([]bool, n)

	// Build the adjacency list
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
	}

	var dfs func(int) bool
	dfs = func(node int) bool {
		visited[node] = true
		recStack[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			} else if recStack[neighbor] {
				// cycle found
				return true
			}
		}

		recStack[node] = false
		return false
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			if dfs(i) {
				return true
			}
		}
	}
	return false
}

func main() {
	n := 4
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 1}, // Cycle: 1 -> 2 -> 3 -> 1
	}

	fmt.Println("Does the directed graph have a cycle?", hasDirectedCycle(n, edges)) // Output: true

	// Cycle exists
	edges = [][]int{{0, 1}, {1, 2}, {2, 0}}                                          // true
	fmt.Println("Does the directed graph have a cycle?", hasDirectedCycle(n, edges)) // Output: true

	// No cycle
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}}                                          // false
	fmt.Println("Does the directed graph have a cycle?", hasDirectedCycle(n, edges)) // Output: false
}

// Want to take it up a notch? Next steps:

// Use this logic to solve Course Schedule (LeetCode 207)
// Implement Topological Sort (it fails if there's a cycle!)
// Or practice cycle detection using Kahn’s Algorithm (BFS-based)
