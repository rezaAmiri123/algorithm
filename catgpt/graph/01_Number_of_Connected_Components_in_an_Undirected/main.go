package main

import "fmt"

// ✅ Problem: Number of Connected Components in an Undirected Graph
// Description:
// You are given n nodes labeled from 0 to n - 1 and a list of undirected edges. Write a function that returns the number of connected components in the graph

// 📥 Input:
// n = 5

// edges = [[0, 1], [1, 2], [3, 4]]

// 📤 Output:
// 2 (because nodes 0-1-2 are connected, and nodes 3-4 are connected)

// 🔧 Approach:
// Use Depth-First Search (DFS) or Union-Find (Disjoint Set Union - DSU). Here's the DFS version in Go:

func countComponents(n int, edges [][]int) int {
	graph := make(map[int][]int)
	visited := make([]bool, n)

	// Build the adjacency list
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		for _, neighbor := range graph[node] {
			dfs(neighbor)
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}
	return count
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Println("Number of connected components:", countComponents(n, edges)) // Output: 2
}

// Would you like to try this problem first and then I can give you variations like:

// Detect cycle in a graph

// Topological sorting

// Shortest path (Dijkstra)

// Union-Find version of the same problem