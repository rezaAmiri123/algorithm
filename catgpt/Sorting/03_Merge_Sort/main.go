package main

import "fmt"

// Great! Merge Sort is another classic sorting algorithm frequently asked in interviews,
// especially because of its predictable O(n log n) time complexity.

// 🔢 Problem Statement:
// Write a Go function to sort an array of integers using Merge Sort.

// ✅ Example Input:
// arr := []int{12, 11, 13, 5, 6, 7}
// 🧠 Expected Output:
// []int{5, 6, 7, 11, 12, 13}

// 🧮 Explanation:
// Merge Sort is a divide-and-conquer algorithm:
// Divide the array into two halves.
// Recursively sort both halves.
// Merge the two sorted halves into one sorted array.

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Divide
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Conquer (merge sorted halves)
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge elements in order
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

// ⏱ Time & Space Complexity:
// Time: O(n log n) in all cases
// Space: O(n) – uses extra space for merging

// ✅ When to Use:
// Merge Sort is especially useful when:
// You need stable sorting.
// Dealing with linked lists.
// You're okay with extra space and want predictable performance.

// Would you like a challenge exercise to implement and compare all three sorts on random arrays?