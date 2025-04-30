package main

import "fmt"

// 🔑 Problem: Two Sum
// Given an array of integers nums and an integer target,
// return the indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.

// ✅ Example:
// Input: nums = [2, 7, 11, 15], target = 9
// Output: [0, 1]
// Because nums[0] + nums[1] == 2 + 7 == 9

// 🧠 Hashing Approach Explanation:
// We use a hash map (map[int]int) to store the numbers we've seen so far and their indices.
// As we iterate through the list, for each number num, we calculate its complement as target - num.
// We check if this complement is already in the map:
// If it is, we return the current index and the index of the complement.
// If not, we add num to the map with its index.
// This gives us O(n) time complexity and O(n) space.

func twoSum(nums []int, target int) []int {
	// map to store number -> index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, ok := numMap[complement]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil // return nil if no solution
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Result:", result)
}

// Would you like to try a slightly more challenging hashing problem next, 
// like finding the first unique character, grouping anagrams, 
// or longest substring without repeating characters?