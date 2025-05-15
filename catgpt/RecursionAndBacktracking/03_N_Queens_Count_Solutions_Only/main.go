package main

import "fmt"

// 🧩 Problem: N-Queens II – Count Solutions Only
// Instead of returning all solutions, just return the total number of valid solutions
// for a given n.

// Example:
// Input: n = 4
// Output: 2
// (There are 2 distinct ways to place 4 queens.)

func totalNQueens(n int) int {
	count := 0

	var cols = make(map[int]bool)
	var diag1 = make(map[int]bool)
	var diag2 = make(map[int]bool)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			count++
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}

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
	return count
}

func main() {
	n := 8
	fmt.Printf("Total number of solutions for %d-Queens: %d\n", n, totalNQueens(n))
}

// 🧠 Key Changes from Full N-Queens
// No board construction.
// We just increment count when a valid configuration is found.
// Much faster and uses less memory – ideal for large n.

// Want to try solving a Sudoku or something more graph-heavy like the Word Search problem next?
