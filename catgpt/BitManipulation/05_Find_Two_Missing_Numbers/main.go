package main

import "fmt"

// 🧠 Problem: Find Two Missing Numbers
// Prompt:
// You are given an array nums containing n - 2 distinct numbers from the range [1, n].
// That means two numbers are missing.
// Return the two missing numbers in O(n) time and O(1) extra space.

// ✅ Example:
// Input:  nums = [1, 2, 4, 6]
// Output: [3, 5]

// 🔍 Explanation:
// This is an advanced XOR trick, similar to the "find two unique numbers" pattern.
// Here's the step-by-step:
// Find the XOR of the two missing numbers (x ^ y):
// XOR all numbers from 1 to n.
// XOR all numbers in the array.
// The result is xor = x ^ y.
// Find a set bit (diffBit) in the XOR result. This bit differs between x and y.
// Partition all numbers (both full range and array) into two groups based on diffBit,
// and XOR within each group.
// You'll isolate x in one group and y in the other.

func findTwoMissingNumbers(nums []int, n int) (int, int) {
	xor := 0

	// Step 1: XOR all numbers from 1 to n
	for i := 0; i <= n; i++ {
		xor ^= i
	}

	// Step 2: XOR all elements in the input array
	for _, num := range nums {
		xor ^= num
	}

	// Step 3: Find rightmost set bit (this distinguishes the two missing numbers)
	diffBit := xor & -xor

	x, y := 0, 0

	// Step 4: Divide numbers in two groups and XOR each group
	for i := 0; i <= n; i++ {
		if i&diffBit == 0 {
			x ^= i
		} else {
			y ^= i
		}
	}

	for _, num := range nums {
		if num&diffBit == 0 {
			x ^= num
		} else {
			y ^= num
		}
	}

	return x, y
}

func main() {
	nums := []int{1, 2, 4, 6}
	n := 6 // total numbers should be 1 through 6
	a, b := findTwoMissingNumbers(nums, n)
	fmt.Println("The two missing numbers are:", a, b)
}

// This problem demonstrates mastery of XOR logic and set-bit partitioning. 
// Want to move on to a bitmask-based subset generation or bitwise dynamic programming?
