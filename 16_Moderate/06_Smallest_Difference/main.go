package main

import (
	"math"
	"slices"
)

// Smallest Difference: Given two arrays of integers, compute the pair of values (one value in each
// 	array) with the smallest (non-negative) difference. Return the difference.
// 	EXAMPLE
// 	Input: {l, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
// 	Output: 3. That is, the pair (11, 8).

// SOLUTION
// Let's start first with a brute force solution.

// Brute Force
// The simple brute force way is to just iterate through all pairs, compute the difference, and compare it to the
// current minimum difference.
func findSmallestDifferences(array1, array2 []int) int {
	if len(array1) == 0 || len(array2) == 0 {
		return -1
	}
	min := math.MaxInt
	for i := 0; i < len(array1); i++ {
		for j := 0; i < len(array2); j++ {
			diff := array1[i] - array2[j]
			if diff < 0 {
				diff *= -1
			}
			if diff < min {
				min = diff
			}
		}
	}
	return min
}

// One minor optimization we could perform from here is to return immediately if we find a difference of
// zero, since this is the smallest difference possible. However, depending on the input, this might actually be
// slower.
// This will only be faster if there's a pair with difference zero early in the list of pairs. But to add this optimiza­
// tion, we need to execute an additional line of code each time. There's a tradeoff here; it's faster for some
// inputs and slower for others. Given that it adds complexity in reading the code, it may be best to leave it out.
// With or without this "optimization;' the algorithm will take O(AB) time.

// Optimal
// A more optimal approach is to sort the arrays. Once the arrays are sorted, we can find the minimum differ­
// ence by iterating through the array.
// Consider the following two arrays:
// A: {l, 2, 11, 15}
// B: {4, 12, 19, 23, 127, 235}
// Try the following approach:
// 1.
// Suppose a pointer a points to the beginning of A and a pointer b points to the beginning of B. The
// current difference between a and b is 3. Store this as the min.
// 2. How can we (potentially) make this difference smaller? Well, the value at b is bigger than the value at a,
// so moving b will only make the difference larger. Therefore, we want to move a.
// 3. Now a points to 2 and b (still) points to 4. This difference is 2, so we should update min. Move a, since
// it is smaller.
// 4. Now a points to 11 and b points to 4. Move b.
// 5. Now a points to 11 and b points to 12. Update min to 1. Move b.
// And so on.
func findSmallestDifferencesImproved(array1, array2 []int) int {
	slices.Sort(array1)
	slices.Sort(array2)
	a:=0
	b:=0
	min := math.MaxInt
	for a<len(array1)&&b<len(array2){
		diff := array1[a]-array2[b]
		if diff <0{
			diff *=-1
		}
		if diff<min{
			min = diff
		}
		if array1[a]<array2[b]{
			a++
		}else{
			b++
		}
	}
	return min
}

// This algorithm takes O(A log A + B log B) time to sort and O(A + B) time to find the minimum
// difference. Therefore, the overall runtime is O(A log A + B log B).
