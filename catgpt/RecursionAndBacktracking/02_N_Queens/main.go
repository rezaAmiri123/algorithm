package main

import (
	"fmt"
	"strings"
)

// 🧩 Problem: N-Queens
// Place n queens on an n x n chessboard so that no two queens attack each other.
// Return all distinct solutions.

// Example for n = 4:
// . Q . .
// . . . Q
// Q . . .
// . . Q .

// . . Q .
// Q . . .
// . . . Q
// . Q . .
// Each solution is a list of strings, where 'Q' marks a queen and '.' an empty space.

// 🧠 Explanation
// Backtracking will help explore every way to place queens:
// Place a queen in a row.
// Move to the next row and try placing another queen in a safe column.
// Backtrack if no position is valid.
// We need to check:
// Column conflicts
// Diagonal conflicts (↘ and ↙)

// We’ll use:
// A cols set for columns
// A diag1 set for top-left to bottom-right diagonals (row - col)
// A diag2 set for top-right to bottom-left diagonals (row + col)

func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([]int, n) // board[i] = column position of queen in row i

	var cols = make(map[int]bool)
	var diag1 = make(map[int]bool)
	var diag2 = make(map[int]bool)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			// Build board from positions
			var solution []string
			for i := 0; i < n; i++ {
				rowStr := strings.Repeat(".", board[i]) + "Q" + strings.Repeat(".", n-board[i]-1)
				solution = append(solution, rowStr)
			}
			res = append(res, solution)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}

			board[row] = col
			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true

			backtrack(row + 1)

			delete(cols, col)
			delete(diag1, row-col)
			delete(diag2, row+col)
		}
	}

	backtrack(0)
	return res
}

func main() {
	n := 4
	solutions := solveNQueens(n)
	fmt.Printf("Total solutions for %d-Queens: %d\n", n, len(solutions))
	for i, sol := range solutions {
		fmt.Printf("Solution %d:\n", i+1)
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println()
	}
}

// 🔍 Key Concepts
// Backtracking tries every queen position in a row.
// Maps (cols, diag1, diag2) track which positions are under attack.
// Build the board only after placing all queens.

// Would you like to try modifying this to return just the number of solutions 
// (like in Leetcode's N-Queens II) or step into Sudoku Solver next?
