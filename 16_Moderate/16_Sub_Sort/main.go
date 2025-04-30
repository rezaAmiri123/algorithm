package main

import "fmt"

// Sub Sort: Given an array of integers, write a method to find indices m and n such that if you sorted
// elements m through n , the entire array would be sorted. Minimize n - m (that is, find the smallest
// such sequence).
// EXAMPLE
// Input: 1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19
// Output: (3, 9)

// SOLUTION
// Before we begin, let's make sure we understand what our answer will look like. If we're looking for just two
// indices, this indicates that some middle section of the array will be sorted, with the start and end of the
// array already being in order.
// Now, let's approach this problem by looking at an example.
// 1, 2, 4, 7, 10, 11, 8, 12, 5, 6, 16, 18, 19
// Our first thought might be to just find the longest increasing subsequence at the beginning and the longest
// increasing subsequence at the end.
// left: 1, 2, 4, 7, 10, 11
// middle: 8, 12
// right: 5, 6, 16, 18, 19

// These subsequences are easy to generate. We just start from the left and the right sides, and work our way
// inward. When an element is out of order, then we have found the end of our increasing/decreasing subse­
// quence.
// In order to solve our problem, though, we would need to be able to sort the middle part of the array and, by
// doing just that, get all the elements in the array in order. Specifically, the following would have to be true:
// /* all items on left are smaller than all items in middle */
// min(middle) > end(left)
// /* all items in middle are smaller than all items in right */
// max(middle) < start(right)
// Or, in other words, for all elements:
// left< middle < right

// In fact, this condition will never be met. The middle section is, by definition, the elements that were out
// of order. That is, it is always the case that left. end > middle. start and middle. end > right.
// start. Thus, you cannot sort the middle to make the entire array sorted.
// But, what we can do is shrink the left and right subsequences until the earlier conditions are met. We need
// the left part to be smaller than all the elements in the middle and right side, and the right part to be bigger
// than all the elements on the left and right side.
// Let min equal min(middle and right side) and max equal max(middle and left side).
// Observe that since the right and left sides are already in sorted order, we only actually need to check their
// start or end point.
// On the left side, we start with the end of the subsequence (value 11, at element 5) and move to the left. The
// value min equals 5. Once we find an element i such that array[ i] < min, we know that we could sort
// the middle and have that part of the array appear in order.

// Then, we do a similar thing on the right side. The value max equals 12.
// right subsequence tvalue 6) and move to the right. We compare the max ot, 2 to 6, then 7, then '\6. When
// reach 16, we know that no elements smaller than 12 could be after it (since it's an increasing subsequence).
// Thus, the middle of the array could now be sorted to make the entire array sorted.
// The following code implements this algorithm.

func findUnsortedSequences(array []int) {
	// find left subsequence
	endLeft := findStartOfLeftSubsequence(array)
	if endLeft >= len(array)-1 {
		return // already sorted
	}
	// find right subsequence
	startRight := findStartOfRightSubsequence(array)

	// get min and max
	maxIndex := endLeft    //max of left side
	minIndex := startRight // min of right side
	for i := endLeft + 1; i < startRight; i++ {
		if array[i] < array[minIndex] {
			minIndex = i
		}
		if array[i] > array[maxIndex] {
			maxIndex = i
		}
	}

	// slide left until less than array[min_index]
	leftIndex := shrinkLeft(array, minIndex, endLeft)

	// slide right until greater than array[max_index]
	rightIndex := shrinkRight(array, maxIndex, startRight)

	fmt.Println(leftIndex, " ", rightIndex)
}

func findStartOfLeftSubsequence(array []int) int {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			return i - 1
		}
	}
	return len(array) - 1
}

func findStartOfRightSubsequence(array []int) int {
	for i := len(array) - 2; i >= 0; i-- {
		if array[i] > array[i+1] {
			return i + 1
		}
	}
	return 0
}

func shrinkLeft(array []int, minIndex, start int) int {
	comp := array[minIndex]
	for i := start - 1; i >= 0; i-- {
		if array[i] <= comp {
			return i + 1
		}
	}
	return 0
}

func shrinkRight(array []int, maxIndex, start int) int {
	comp := array[maxIndex]
	for i := start; i < len(array); i++ {
		if array[i] >= comp {
			return i - 1
		}
	}
	return len(array) - 1
}

// Note the use of other methods in this solution. Although we could have jammed it all into one method, it
// would have made the code a lot harder to understand, maintain, and test. In your interview coding, you
// should prioritize these aspects.
