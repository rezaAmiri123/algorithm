package main

import "fmt"

// 🔑 Problem: Longest Consecutive Sequence
// Given an unsorted array of integers nums,
// return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.

// Input: nums = [100, 4, 200, 1, 3, 2]
// Output: 4
// Explanation: The longest consecutive sequence is [1, 2, 3, 4]

// 🧠 Hashing Strategy (O(n)):
// Put all numbers in a map[int]bool (or use map[int]struct{} for lighter memory).
// Iterate through each number:
// Only start counting when num - 1 does not exist (i.e., it's the start of a sequence).
// From there, count how long the streak continues (num + 1, num + 2, ...).
// Track the max streak length.

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0

	for num := range numSet {
		// Check if it's the start of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Longest consecutive sequence length:", longestConsecutive(nums))
}

// 🔥 This is a true O(n) solution — perfect for interviews where brute force isn't acceptable.
// Wanna level it up and return the actual sequence instead of just the length? 
// Or try something like Top K Frequent Elements next?
