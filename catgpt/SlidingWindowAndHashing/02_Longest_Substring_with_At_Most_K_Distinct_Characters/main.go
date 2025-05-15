package main

import "fmt"

// Great! Here's a harder variant of the sliding window + hashing problem:

// 🔍 Problem: Longest Substring with At Most K Distinct Characters
// Given a string s and an integer k, return the length of the longest substring
// that contains at most k distinct characters.

// Example:
// Input: s = "eceba", k = 2
// Output: 3
// Explanation: The substring is "ece" with length 3.

// 💡 Explanation
// This builds on the previous problem but now allows repeated characters,
// as long as the number of distinct characters is ≤ k.

// Steps:
// Use a map[byte]int to store character counts in the current window.
// Use two pointers: start and end (sliding window).
// Expand end, updating the map.
// If the map size exceeds k, shrink from start until it’s valid again.
// Track the max window length along the way.

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	charCount := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		charCount[s[end]]++

		for len(charCount) > k {
			charCount[s[start]]--
			if charCount[s[start]] == 0 {
				delete(charCount, s[start])
			}
			start++
		}

		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}

	return maxLen
}

func main() {
	s := "eceba"
	k := 2
	fmt.Println("Length of longest substring with at most", k, "distinct characters:", lengthOfLongestSubstringKDistinct(s, k))
}

// 🧠 Why it's harder:
// Requires handling frequency counts.
// Requires understanding of map size control.
// More nuanced edge cases (k=0, string length 0, etc).
// Would you like a follow-up version where instead of characters, you work with arrays 
// and integers?
