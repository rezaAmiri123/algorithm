package main

import (
	"fmt"
	"sort"
)

// Great! Let's implement the Knight’s Tour using Warnsdorff’s Rule,
// which is a famous heuristic that makes the brute-force search vastly more efficient.

// 🧠 What is Warnsdorff’s Rule?
// It chooses the next move that has the fewest onward moves (i.e., least degree of freedom).
// This reduces the chance of early dead-ends by guiding the knight into tighter spaces earlier.

// Instead of brute-forcing all 8 directions equally,
// we sort them by how many onward valid moves each one leads to.

var moves = [8][2]int{
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
}

type Move struct {
	x, y, degree int
}

func isValid(x, y, n int, board [][]int) bool {
	return x >= 0 && y >= 0 && x < n && y < n && board[x][y] == -1
}

// Count onward moves from (x, y)
func countOnwardMoves(x, y, n int, board [][]int) int {
	count := 0
	for _, m := range moves {
		nx, ny := x+m[0], y+m[1]
		if isValid(nx, ny, n, board) {
			count++
		}
	}
	return count
}

func knightTourWarnsdorff(x, y, moveNum, n int, board [][]int) bool {
	if moveNum == n*n {
		return true
	}

	var nextMoves []Move
	for _, m := range moves {
		nx, ny := x+m[0], y+m[1]
		if isValid(nx, ny, n, board) {
			degree := countOnwardMoves(nx, ny, n, board)
			nextMoves = append(nextMoves, Move{nx, ny, degree})
		}
	}

	// Sort by Warnsdorff's Rule: fewest onward moves first
	sort.Slice(nextMoves, func(i, j int) bool {
		return nextMoves[i].degree < nextMoves[j].degree
	})

	for _, move := range nextMoves {
		board[move.x][move.y] = moveNum
		if knightTourWarnsdorff(move.x, move.y, moveNum+1, n, board) {
			return true
		}
		board[move.x][move.y] = -1 // backtrack
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

	if !knightTourWarnsdorff(0, 0, 1, n, board) {
		fmt.Println("No solution exists.")
		return nil
	}
	return board
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}
}

func main() {
	n := 8
	board := solveKnightTour(n)
	if board != nil {
		fmt.Println("Knight's Tour using Warnsdorff’s Rule:")
		printBoard(board)
	}
}
// 🧪 Try It With:
// n := 8 – Classic chessboard size.
// n := 6 or 7 – Easier for quick validation.

// ✅ Advantages Over Brute-Force
// Much faster, especially on larger boards.
// Almost always finds a solution without full backtracking.
// Elegant and efficient with zero or minimal recursion depth.

// Would you like to try modifying this to generate all possible tours (non-heuristic), 
// or jump to a maze solver, or a constraint puzzle like Kakuro or KenKen next?
