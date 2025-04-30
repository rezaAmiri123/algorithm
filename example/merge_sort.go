package main

import "fmt"

// merge combines two sorted slices into a single sorted slice.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// Append remaining elements
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// mergeSort sorts a slice of integers using the merge sort algorithm.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	data := []int{9, 2, 7, 1, 5, 6, 3, 8, 4}
	fmt.Println("Unsorted:", data)
	sortedData := mergeSort(data)
	fmt.Println("Sorted:  ", sortedData)
}
