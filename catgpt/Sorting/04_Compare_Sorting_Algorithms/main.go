package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 💪 Challenge: Compare Sorting Algorithms in Go
// 🔧 Task:
// Implement Bubble Sort, Quick Sort, and Merge Sort. Generate a random slice of integers
// and sort it using each algorithm. Compare their performance using Go's time package.

// 🔁 What You'll Learn:
// How different algorithms behave on the same input.
// How Go handles performance benchmarking.
// Solidifying your understanding of sorting logic.

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

// Generate random integers
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(10000)
	}
	return slice
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := generateRandomSlice(10000)

	start := time.Now()
	bubbleSort(arr)
	fmt.Println("Bubble Sort took:", time.Since(start))

	start = time.Now()
	quickSort(arr)
	fmt.Println("Quick Sort took:", time.Since(start))

	start = time.Now()
	mergeSort(arr)
	fmt.Println("Merge Sort took:", time.Since(start))
}

// ⚠️ Bubble Sort is dramatically slower on large inputs. 
// You might want to reduce the slice size to 1000 or less to avoid long waits.

// Would you like a version with sorting validation 
// and optional slice size from command line input?
