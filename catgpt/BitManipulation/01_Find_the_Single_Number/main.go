package main

import "fmt"

// 🧠 Problem: Find the Single Number
// Prompt:
// Given a non-empty array of integers nums, every element appears twice except for one.
// Find that single one.
// You must implement a solution with a linear runtime complexity
// and use only constant extra space.

// ✅ Example:
// Input:  nums = [4,1,2,1,2]
// Output: 4

// 🔍 Explanation:
// This is a classic use-case for the XOR bitwise operation. The idea is:
// XOR (^) of a number with itself is 0: a ^ a = 0
// XOR of a number with 0 is the number: a ^ 0 = a
// XOR is commutative and associative, meaning the order doesn't matter.
// So, if every number appears twice except one,
// XOR-ing all of them will cancel out the duplicates and leave the single number.

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println("Single number is:", singleNumber(nums))
}

// Would you like a harder problem next, such as detecting the rightmost set bit 
// or finding two non-repeating elements?
