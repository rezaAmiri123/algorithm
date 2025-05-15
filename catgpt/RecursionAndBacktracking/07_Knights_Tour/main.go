package main

import "fmt"

// Awesome! Let's tackle the Knight’s Tour, a fascinating backtracking + graph traversal problem.
// It’s less common than Sudoku but really shows off deep algorithmic thinking
// and recursion control.

// 🧩 Problem: Knight’s Tour
// Given an N x N chessboard,
// find one valid tour where the knight visits every square exactly once.
// The knight moves in an "L" shape: two squares in one direction, then one square perpendicular.

// We’ll solve the classic version:
// Start at (0, 0)
// Fill the board with the order in which the knight visits each square
// Return one solution

// 📐 Example Output for N = 5 (numbers = step index):
// 0 17  8 25 20
// 9 24 19 16  7
// 18  1 10 21 26
// 23 12  3  6 15
// 2 27 22 13  4

var (
	moves = [8][2]int{
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	}
)

func isValid(x, y, n int, board [][]int) bool {
	return x >= 0 && y >= 0 && x < n && y < n && board[x][y] == -1
}

func knightTour(x, y, moveNum, n int, board [][]int) bool {
	if moveNum == n*n {
		return true
	}

	for _, move := range moves {
		nextX, nextY := x+move[0], y+move[1]
		if isValid(nextX, nextY, n, board) {
			board[nextX][nextY] = moveNum
			if knightTour(nextX, nextY, moveNum+1, n, board) {
				return true
			}
			board[nextX][nextY] = -1 // backtrack
		}
	}
	return false
}

func solveKnightTour(n int) [][]int {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
		for j := range board[i] {
			board[i][j] = -1
		}
	}
	board[0][0] = 0

	if !knightTour(0, 0, 1, n, board) {
		fmt.Println("No solution exists.")
		return nil
	}

	return board
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%2d ", cell)
		}
		fmt.Println()
	}
}

func main() {
	n := 5
	board := solveKnightTour(n)
	if board != nil {
		fmt.Println("Knight's Tour:")
		printBoard(board)
	}
}

// 🧠 How It Works
// knightTour recursively tries all knight moves from current position.
// moveNum tracks the step count.
// If moveNum == n*n, the board is complete.
// Backtracking resets invalid moves (board[x][y] = -1).
// Only one solution is returned (you could modify to find all).

// This version is brute-force. Want to upgrade it with Warnsdorff’s Rule 
// (a heuristic for faster solutions), 
// or try a completely different challenge like maze solving 
// or backtracking with constraints (like solving puzzles)?