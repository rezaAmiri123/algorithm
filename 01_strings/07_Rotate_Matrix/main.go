package main

import "fmt"

// Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

// SOLUTION
// Because we're rotating the matrix by 90 degrees, the easiest way to do this is to implement the rotation in
// layers. We perform a circular rotation on each layer, moving the top edge to the right edge, the right edge
// to the bottom edge, the bottom edge to the left edge, and the left edge to the top edge.

// How do we perform this four-way edge swap? One option is to copy the top edge to an array, and then
// move the left to the top, the bottom to the left, and so on. This requires O(N) memory, which is actually
// unnecessary.
// A better way to do this is to implement the swap index by index. In this case, we do the following:
// for i = 0 to n
// temp= top[i];
// top[i] = left[i]
// left[i] = bottom[i]
// bottom[i] = right[i]
// right[i] = temp

// We perform such a swap on each layer, starting from the outermost layer and working our way inwards.
// (Alternatively, we could start from the inner layer and work outwards.)
// The code for this algorithm is below.

func rotate(matrix  [][]int)bool{
	if len(matrix) == 0||len(matrix[0])!=len(matrix){
		return false
	}
	n := len(matrix)// n=3
	for layer := 0;layer <n/2;layer++{
		first := layer //layer =0, first = 0-->layer =1, first = 1
		last := n-1-layer // last = 2 --> last = 1
		for i:=first;i<last;i++{
			offset := i-first // offset= 0, first = 0, -->offset= 1, first = 0,

			top := matrix[first][i] // save top =(0,0)-->top =(0,1)
			
			// left -> top
			matrix[first][i] = matrix[last-offset][first] 
			// m(0,0) = m(2,0)-->m(0,1) = m(1,0)

			// bottom -> left
			matrix[last-offset][first]=matrix[last][last-offset] 
			// m(2,0) = m(2,2)-->m(1,0) = m(2,1)

			// right -> bottom
			matrix[last][last-offset]=matrix[i][last] 
			// m(2,2) = m(0,2)-->m(2,1) = m(1,2)

			// top -> right
			matrix[i][last]=top // right <- saved top 
			// m(0,2) = top --> m(1,2) = top
		}
	}
	fmt.Println(matrix)
	return true
}

// This algorithm is O( N2 ), which is the best we can do since any algorithm must touch all N2 elements.

func main(){
	input := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println("rotate: ",rotate(input))

}