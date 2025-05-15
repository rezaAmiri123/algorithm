package main

import "fmt"

// 🧠 Problem: Find the Missing Number
// Prompt 1 (Classic Version):
// Given an array nums containing n distinct numbers in the range [0, n],
// return the one number that is missing from the array.

// You must solve it in O(n) time using O(1) extra space.
//
// ✅ Example:
// Input:  nums = [3, 0, 1]
// Output: 2

// 🔍 Bitwise XOR Explanation:
// You can find the missing number using XOR with this logic:
// XOR all numbers from 0 to n.
// XOR all elements in the array.
// XOR the two results — duplicates cancel out, and the missing number remains.

func missingNumber(nums []int) int {
	n := len(nums)
	xor := 0

	for i := 0; i <= n; i++ {
		xor ^= i
	}

	for _, num := range nums {
		xor ^= num
	}

	return xor
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println("Missing number is:", missingNumber(nums))
}

// 🧠 Bonus Prompt: Find Two Missing Numbers (Harder!)
// If you're up for even more, 
// I can walk you through finding two missing numbers in a range using bitwise tricks — 
// let me know!
