package main

// Great! Here's a more advanced and flexible version of the sorting comparison program.
// This version:
// Takes the slice size from command line input
// Validates that all sorting algorithms return the same result
// Uses a default size (1000) if no argument is given

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Bubble Sort
func bubbleSort(arr []int) []int {
	n := len(arr)
	copyArr := make([]int, n)
	copy(copyArr, arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if copyArr[j] > copyArr[j+1] {
				copyArr[j], copyArr[j+1] = copyArr[j+1], copyArr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return copyArr
}

// Quick Sort
func quickSort(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	qs(copyArr, 0, len(copyArr)-1)
	return copyArr
}

func qs(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		qs(arr, low, pi-1)
		qs(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// Utilities
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100000)
	}
	return slice
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get size from CLI or use default
	size := 1000
	if len(os.Args) > 1 {
		if s, err := strconv.Atoi(os.Args[1]); err == nil && s > 0 {
			size = s
		} else {
			fmt.Println("Invalid input. Using default size 1000.")
		}
	}
	fmt.Printf("Sorting %d random integers...\n", size)

	original := generateRandomSlice(size)

	// Bubble Sort
	start := time.Now()
	bubble := bubbleSort(original)
	fmt.Printf("Bubble Sort took: %v\n", time.Since(start))

	// Quick Sort
	start = time.Now()
	quick := quickSort(original)
	fmt.Printf("Quick Sort took: %v\n", time.Since(start))

	// Merge Sort
	start = time.Now()
	merge := mergeSort(original)
	fmt.Printf("Merge Sort took: %v\n", time.Since(start))

	// Validation
	if isEqual(bubble, quick) && isEqual(quick, merge) {
		fmt.Println("✅ All sorting algorithms produced the same result.")
	} else {
		fmt.Println("❌ Sorting results differ!")
	}
}
// go run sort_compare.go           # Defaults to 1000 elements
// go run sort_compare.go 5000     # Custom number of elements
// ⚠️ Bubble Sort is slow. Avoid sizes > 2000 unless you're testing Quick/Merge only.

// Would you like to log the sorted output or save it to a file for debugging?