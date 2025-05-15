package main

import "fmt"

// Perfect choice — Quick Sort is a very common algorithm in interviews due to its efficiency
// and elegance.

// 🔢 Problem Statement:
// Write a Go function to sort an array of integers using Quick Sort.

// ✅ Example Input:
// arr := []int{10, 7, 8, 9, 1, 5}
// 🧠 Expected Output:
// []int{1, 5, 7, 8, 9, 10}

// 🧮 Explanation:
// Quick Sort is a divide-and-conquer algorithm:
// Pick a pivot element.
// Partition the array such that:
// All elements smaller than the pivot go to the left.
// All elements greater than the pivot go to the right.
// Recursively apply the above steps to left and right subarrays.

func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pi := Partition(arr, low, high)

		// Recursively sort elements before and after partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose the last element as pivot
	i := low - 1       // Index of smaller element

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap
		}
	}
	// Place pivot in correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}

// ⏱ Time Complexity:
// Best / Average case: O(n log n)

// Worst case: O(n²) – occurs when the pivot is the smallest or largest element each time 
// (can be improved with randomization or median-of-three)

// Would you like to try implementing Merge Sort next 
// or perhaps practice some interview-style sorting questions?
