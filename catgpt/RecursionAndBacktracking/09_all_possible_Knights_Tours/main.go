package main

import "fmt"

// Awesome — now we're going full brute-force:
// let's find all possible Knight’s Tours on an N x N board!

// ⚠️ Important Notes
// The number of possible knight's tours grows explosively.
// For n = 5, there are thousands of solutions.
// For n = 6+, the number becomes impractical for full enumeration without advanced optimizations.
// We’ll use recursive backtracking without Warnsdorff’s rule.
// This version will count all valid tours, and you can optionally print them.

var moves = [8][2]int{
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
}

func isValid(x, y, n int, board [][]int) bool {
	return x >= 0 && y >= 0 && x < n && y < n && board[x][y] == -1
}

func knightTourAll(x, y, moveNum, n int, board [][]int, count *int) {
	if moveNum == n*n {
		*count++
		// Uncomment to print each solution:
		// fmt.Println("Tour", *count)
		// printBoard(board)
		return
	}

	for _, move := range moves {
		nx, ny := x+move[0], y+move[1]
		if isValid(nx, ny, n, board) {
			board[nx][ny] = moveNum
			knightTourAll(nx, ny, moveNum+1, n, board, count)
			board[nx][ny] = -1 // backtrack
		}
	}
}

func countKnightTours(n int) int {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
		for j := range board[i] {
			board[i][j] = -1
		}
	}
	board[0][0] = 0
	count := 0
	knightTourAll(0, 0, 1, n, board, &count)
	return count
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	n := 5
	fmt.Printf("Counting all knight's tours on %dx%d board...\n", n, n)
	count := countKnightTours(n)
	fmt.Printf("Total knight's tours: %d\n", count)
}

// 🧠 Summary
// Brute-force recursion explores every valid path.
// Backtracking ensures no square is visited twice.
// You can uncomment printBoard() to see each solution (will flood output for n ≥ 5).

// 🚀 Optimization Tip
// You can speed it up by:
// Symmetry pruning: e.g., only start from upper-left quadrant and multiply by symmetry factor.
// Using Warnsdorff’s rule to guide initial steps even in exhaustive search.
// Using parallel execution (Goroutines) for different starting squares.

// Would you like me to add symmetry pruning or Goroutine-based parallelism to count faster?
