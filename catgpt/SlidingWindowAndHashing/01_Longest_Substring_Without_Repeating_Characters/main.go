package main

import "fmt"

// Great choice! Sliding Window combined with Hashing is a common
// and powerful pattern for solving problems involving substrings, arrays, or sequences.
// Here's a classic problem that uses both:

// 🔍 Problem: Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// 💡 Explanation
// We use a sliding window to maintain a window of non-repeating characters
// and a hash map to track the last seen index of characters.

// Steps:
// Use a map to keep track of characters and their latest positions.
// Maintain a start pointer for the beginning of the window.
// Iterate through the string with index end:
// If the character has been seen and is inside the window (map[c] >= start), move the start
// to map[c] + 1.
// Update the character's latest position in the map.
// Track the max window length.

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		c := s[end]
		if idx, found := charIndex[c]; found && idx >= start {
			start = idx + 1
		}
		charIndex[c] = end
		windowLength := end - start + 1
		if windowLength > maxLength {
			maxLength = windowLength
		}
	}
	return maxLength
}

func main() {
	s := "abcabcbb"
	fmt.Println("Length of longest substring without repeating characters:", lengthOfLongestSubstring(s))
}

// 🧠 Why this is good for interviews:
// Tests understanding of maps and strings in Go.
// Requires handling of indices and edge cases.
// Efficient: O(n) time and O(k) space, where k is the number of distinct characters.
// Would you like to try a harder version or a variation of this problem?
