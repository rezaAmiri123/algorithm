package main

import "fmt"

// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
// column are set to 0.

// SOLUTION
// At first glance, this problem seems easy: just iterate through the matrix and every time we see a cell with
// value zero, set its row and column to 0. There's one problem with that solution though: when we come
// across other cells in that row or column, we'll see the zeros and change their row and column to zero. Pretty
// soon, our entire matrix will be set to zeros.
// One way around this is to keep a second matrix which flags the zero locations. We would then do a second
// pass through the matrix to set the zeros.This would take O(MN) space.

// Do we really need O(MN) space? No. Since we're going to set the entire row and column to zero, we don't
// need to track that it was exactly cell [ 2] [ 4] (row 2, column 4). We only need to know that row 2 has a
// zero somewhere, and column 4 has a zero somewhere.We'll set the entire row and column to zero anyway,
// so why would we care to keep track of the exact location of the zero?

// The code below implements this algorithm. We use two arrays to keep track of all the rows with zeros and all
// the columns with zeros. We then nullify rows and columns based on the values in these arrays.

func setZeros(matrix [][]int) {
	rows := make([]bool, len(matrix))
	column := make([]bool, len(matrix[0]))
	// Store the row and column index with value 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				column[j] = true
			}
		}
	}
	// Nullify rows
	for i := 0; i < len(rows); i++ {
		if rows[i] {
			// nullify rows
			for j:=0;j<len(matrix[0]);j++{
				matrix[i][j] = 0
			}
		}
	}
	// Nullify columns
	for j := 0; j < len(column); j++ {
		if column[j] {
			// nullify column
			for i:=0;i<len(matrix);i++{
				matrix[i][j]=0
			} 
		}
	}
}

func main(){
	matrix := [][]int{
		{1,2,3},
		{4,0,6},
		{7,8,9},
	}
	fmt.Println(matrix)
	setZeros(matrix)
	fmt.Println(matrix)
}
