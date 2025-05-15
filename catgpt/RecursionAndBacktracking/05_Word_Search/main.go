package main

import "fmt"

// Great choice! Word Search is a classic recursion + backtracking problem,
// and it’s often asked in interviews to test your understanding of 2D grid traversal
// and state management.

// 🧩 Problem: Word Search
// Given a 2D board and a word, return true if the word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells,
// where "adjacent" cells are horizontally or vertically neighboring.
// The same letter cell may not be used more than once.

// board := [][]byte{
// 	{'A','B','C','E'},
// 	{'S','F','C','S'},
// 	{'A','D','E','E'},
// }
// word := "ABCCED"
// Output: true

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	var backtrack func(r, c, index int) bool
	backtrack = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[index] {
			return false
		}

		// Mark cell as visited by replacing with '#'
		temp := board[r][c]
		board[r][c] = '#'

		// Explore all directions: up, down, left, right
		found := backtrack(r+1, c, index+1) ||
			backtrack(r-1, c, index+1) ||
			backtrack(r, c+1, index+1) ||
			backtrack(r, c-1, index+1)

		// Backtrack: restore original value
		board[r][c] = temp

		return found
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == word[0] && backtrack(r, c, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Printf("Does the word \"%s\" exist in the board? %v\n", word, exist(board, word))
}

// 🧠 How It Works
// Start DFS at every cell matching the first character of the word.
// Recursively try all four directions.
// Use a temporary marker ('#') to avoid revisiting cells.
// Restore the cell after the recursive call (backtrack).

// Want to try extending this to find all words from a dictionary in the board (like in Word Search II)?
