package main

import "fmt"

// Great! Practicing sorting algorithms in Go is excellent prep for technical interviews.
// Let's go through a classic example: Implementing and understanding the Bubble Sort algorithm.

// 🔢 Problem Statement:
// Write a function in Go to sort an array of integers using Bubble Sort.

// ✅ Example Input:
// arr := []int{5, 2, 9, 1, 5, 6}
// 🧠 Expected Output:
// []int{1, 2, 5, 5, 6, 9}

// 🧮 Explanation:
// Bubble Sort works by repeatedly swapping the adjacent elements if they are in the wrong order.
// It continues doing this until the array is sorted.
// Each pass through the array "bubbles" the largest unsorted element to its correct position.

func bubbleSort(arr []int) []int {
	n := len(arr)
	// Outer loop for each pass
	for i := 0; i < n-1; i++ {
		swapped := false
		// Inner loop for comparing adjacent elements
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// Optimization: stop if no swaps in the inner loop
		if !swapped {
			break
		}
	}
	return arr
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	sortedArr := bubbleSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

// 🚀 Tips:
// Understand time complexity: Bubble Sort is O(n²) in worst and average case, 
// and O(n) in best case (when the array is already sorted).

// Practice implementing other sorts like Selection Sort, Insertion Sort, Merge Sort, 
// Quick Sort for a well-rounded understanding.

// Would you like the next example to be a more efficient algorithm like Quick Sort 
// or Merge Sort?
