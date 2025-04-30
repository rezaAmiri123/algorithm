package main

// Sorted Matrix Search: Given an M x N matrix in which each row and each column is sorted in
// ascending order, write a method to find an element.

// SOLUTION
// We can approach this in two ways: a more naive solution that only takes advantage of part of the sorting,
// and a more optimal way that takes advantage of both parts of the sorting.
// Solution #1: Naive Solution
// As a first approach, we can do binary search on every row to find the element. This algorithm will be O(M
// log( N)), since there are M rows and it takes 0( log( N)) time to search each one. This is a good approach
// to mention to your interviewer before you proceed with generating a better algorithm.
// To develop an algorithm, let's start with a simple example.
// 15 20 40  85
// 20 35 80  95
// 30 55 95  105
// 40 80 100 120

// Suppose we are searching for the element 55. How can we identify where it is?
// If we look at the start of a row or the start of a column, we can start to deduce the location. If the start of a
// column is greater than 55, we know that 55 can't be in that column, since the start of the column is always
// the minimum element. Additionally, we know that 55 can't be in any columns on the right, since the first
// element of each column must increase in size from left to right. Therefore, if the start of the column is
// greater than the element x that we are searching for, we know that we need to move further to the left.

// For rows, we use identical logic. If the start of a row is bigger than x, we know we need to move upwards.
// Observe that we can also make a similar conclusion by looking at the ends of columns or rows. If the end
// of a column or row is less than x, then we know that we must move down (for rows) or to the right (for
// columns) to find x. Thisis because the end is always the maximum element.
// We can bring these observations together into a solution.1he observations are the following:
// If the start of a column is greater than x, then xis to the left of the column.
// If the end of a column is less than x, then x is to the right of the column.
// If the start of a row is greater than x, then x is above that row.
// If the end of a row is less than x, then x is below that row.
// We can begin in any number of places, but let's begin with looking at the starts of columns.

// We need to start with the greatest column and work our way to the left. This means that our first element
// for comparison is array[0] [ c-1], where c is the number of columns. By comparing the start of columns
// to x (which is 55), we'll find that x must be in columns 0, 1, or 2. We will have stopped at array[0] [2].
// This element may not be the end of a row in the full matrix, but it is an end of a row of a submatrix. The
// same conditions apply. The value at array[0] [2], which is 40, is less than 55, so we know we can move
// downwards.
// We now have a submatrix to consider that looks like the following (the gray squares have been eliminated).

// We can repeatedly apply these conditions to search for 55. Note that the only conditions we actually use
// are conditions 1 and 4.
// The code below implements this elimination algorithm.

func findElementNaive(matrix [][]int, elem int) bool {
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == elem {
			return true
		} else if matrix[row][col] > elem {
			col--
		} else {
			row++
		}
	}
	return false
}

// Alternatively, we can apply a solution that more directly looks like binary search. The code is considerably
// more complicated, but it applies many of the same learnings.

// Solution #2: Binary Search
// Let's again look at a simple example.

// We want to be able to leverage the sorting property to more efficiently find an element. So, we might ask
// ourselves, what does the unique ordering property of this matrix imply about where an element might be
// located?
// We are told that every row and column is sorted. This means that element a [ i] [ j] will be greater than
// the elements in row i between columns O and j - 1 and the elements in column j between rows O and
// Or, in other words:
// a[i][0] <= a[i][l] <= ... <= a[i][j-1] <= a[i][j]
// a [0][j] <= a[l][j] <= ... <= a[i-l][j] <= a[i][j]
// Looking at this visually, the dark gray element below is bigger than all the light gray elements.

// Observe that since the diagonal is sorted, we can efficiently search it using binary search.
// The code below implements this algorithm
type Coorditate struct {
	row, column int
}

func (c *Coorditate) inbounds(matrix [][]int) bool {
	return c.row >= 0 && c.column >= 0 &&
		c.row < len(matrix) && c.column <= len(matrix[0])
}

func (c *Coorditate) isBefore(p *Coorditate) bool {
	return c.row <= p.row && c.column <= p.column
}
func (c *Coorditate) setToAverage(min, max *Coorditate) {
	c.row = (max.row + min.row) / 2
	c.column = (max.column + min.column) / 2
}
func (c *Coorditate) clone() *Coorditate {
	return &Coorditate{
		row:    c.row,
		column: c.column,
	}
}
func findElement(matrix [][]int, origin, dest *Coorditate, x int) *Coorditate {
	if !origin.inbounds(matrix) || !dest.inbounds(matrix) {
		return nil
	}
	if matrix[origin.row][origin.column] == x {
		return origin
	} else if !origin.isBefore(dest) {
		return nil
	}
	// 	Set start to start of diagonal and end to the end of the diagonal. Since the
	// * grid may not be square, the end of the diagonal may not equal dest.
	start := origin.clone()
	diagDist := dest.row - origin.row
	column := dest.column - origin.column
	if column < diagDist {
		diagDist = column
	}
	p := &Coorditate{row: 0, column: 0}
	end := &Coorditate{row: start.row + diagDist, column: start.column + diagDist}

	// Do binary search on the diagonal, looking for the first element> x
	for start.isBefore(end) {
		p.setToAverage(start, end)
		if x > matrix[p.row][p.column] {
			start.row = p.row + 1
			start.column = p.column + 1
		} else {
			end.row = p.row - 1
			end.column = p.column - 1
		}
	}
	// Split the grid into quadrants. Search the bottom left and the top right.
	return partitionAndSearch(matrix, origin, dest, start, x)
}

func FindElement(matrix [][]int, x int) *Coorditate {
	origin := &Coorditate{row: 0, column: 0}
	dest := &Coorditate{row: len(matrix) - 1, column: len(matrix[0]) - 1}
	return findElement(matrix, origin, dest, x)
}

func partitionAndSearch(matrix [][]int, orign, dest, pivot *Coorditate, x int) *Coorditate {
	lowerLeftOrigin := &Coorditate{row: pivot.row, column: orign.column}
	lowerLeftDest := &Coorditate{row: dest.row, column: pivot.column - 1}
	upperRightOrigin := &Coorditate{row: orign.row, column: pivot.column}
	upperRightDest := &Coorditate{row: pivot.row - 1, column: dest.column}

	lowerLeft := findElement(matrix, lowerLeftOrigin, lowerLeftDest, x)
	if lowerLeft == nil {
		return findElement(matrix, upperRightOrigin, upperRightDest, x)
	}

	return lowerLeft
}

// If you read all this code and thought, "there's no way I could do all this in an interview!" you're probably
// right. You couldn't. But, your performance on any problem is evaluated compared to other candidates on
// the same problem. So while you couldn't implement all this, neither could they. You are at no disadvantage
// when you get a tricky problem like this.
// You help yourself out a bit by separating code out into other methods. For example, by pulling
// partitionAndSearch out into its own method, you will have an easier time outlining key aspects of the
// code. You can then come back to fill in the body for partitionAndSearch if you have time.
