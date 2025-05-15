package main

import "fmt"

// 🧠 Problem: Find the Two Non-Repeating Elements
// Prompt:
// Given an array of integers where every element appears exactly twice except for two elements
// that appear only once, return the two unique elements.
// The solution should run in linear time and use constant extra space.

// ✅ Example:
// Input:  nums = [1, 2, 1, 3, 2, 5]
// Output: [3, 5]
// (The order doesn't matter.)

// 🔍 Explanation:
// This problem is an extension of the single-number problem and uses clever XOR logic:
// XOR all elements → this gives xor = a ^ b (where a and b are the two unique numbers).
// Find any set bit in xor (usually the rightmost set bit) → this bit differs between a and b.
// Use this bit to divide the array into two groups:
// One where that bit is set.
// One where that bit is not set.
// XOR elements in each group → you'll get a and b.

func singleNumbers(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	// Get rightmost set bit (differentiating bit between the two unique numbers)
	diffBit := xor & -xor

	a, b := 0, 0
	for _, num := range nums {
		if num&diffBit == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	result := singleNumbers(nums)
	fmt.Println("The two non-repeating numbers are:", result)
}

// This problem tests multiple bitwise concepts: XOR logic, masking with set bits, 
// and problem decomposition.
// Would you like a challenge problem next, 
// like counting bits from 0 to n or finding missing numbers using bitwise math?
