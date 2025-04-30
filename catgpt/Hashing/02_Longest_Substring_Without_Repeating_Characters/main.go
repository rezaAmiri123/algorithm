package main

import "fmt"

// 🔑 Problem: Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.

// ✅ Example:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// 🧠 Hashing + Sliding Window Explanation:
// We use a sliding window and a hash map to track characters and their latest index.
// Steps:
// Start with two pointers: left and right. Both start at 0.
// Use a map charIndex to store the last seen index of each character.
// As you expand the window (right moves), if the character has been seen before
// and its index is ≥ left, move left just past that character.
// Update maxLen with the size of the current window (right - left + 1).
// Repeat until the end of the string.

// Time complexity: O(n)
// Space complexity: O(min(n, m)) where m is the size of the charset.

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, found := charIndex[s[right]]; found && idx >= left {
			left = idx + 1
		}
		charIndex[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println("Length of longest substring without repeating characters:", lengthOfLongestSubstring(s))
}

// Want to go one level deeper and modify this to return the actual substring too? 
// Or try another popular one like group anagrams or top K frequent elements?