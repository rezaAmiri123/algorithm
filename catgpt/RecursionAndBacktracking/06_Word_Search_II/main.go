package main

import "fmt"

// Excellent! Now we're diving into Word Search II,
// a tougher and more interview-relevant version of the previous problem.

// 🧩 Problem: Word Search II
// Given a 2D board and a list of words from the dictionary,
// return all words that can be formed by letters of sequentially adjacent cells.

// You may use each cell only once per word.
// board := [][]byte{
// 	{'o','a','a','n'},
// 	{'e','t','a','e'},
// 	{'i','h','k','r'},
// 	{'i','f','l','v'},
// }
// words := []string{"oath","pea","eat","rain"}
// Output: ["eat","oath"]

// 💡 Optimization Tip: Use a Trie (Prefix Tree)
// Instead of checking every word separately (which is slow),
// we insert all words into a Trie and explore the board once, pruning invalid paths early.

type TrieNode struct {
	children map[byte]*TrieNode
	word     string // non-empty when it's the end of a word
}

func buildTrie(words []string) *TrieNode {
	root := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			c := word[i]
			if _, ok := node.children[c]; !ok {
				node.children[c] = &TrieNode{children: make(map[byte]*TrieNode)}
			}
			node = node.children[c]
		}
		node.word = word
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	root := buildTrie(words)
	rows, cols := len(board), len(board[0])
	var result []string

	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		letter := board[r][c]
		child, exists := node.children[letter]
		if !exists {
			return
		}

		if child.word != "" {
			result = append(result, child.word)
			child.word = "" // Avoid duplicates
		}

		board[r][c] = '#' // Mark visited

		directions := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && board[nr][nc] != '#' {
				dfs(nr, nc, child)
			}
		}

		board[r][c] = letter // Backtrack

		// Optimization: prune the trie
		if len(child.children) == 0 {
			delete(node.children, letter)
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}

	return result
}

func main() {
	board := [][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words := []string{"oath","pea","eat","rain"}
	foundWords := findWords(board, words)
	fmt.Println("Found words:", foundWords)
}

// 🧠 How It Works
// A Trie stores the word list for quick prefix matching.
// DFS explores from every cell.
// If a full word is found (node.word != ""), it's added to results.
// Trie pruning removes unnecessary nodes for faster performance.

// Would you like to try a more advanced backtracking problem like Sudoku Solver, 
// or a graph-focused one like Knight’s Tour?
