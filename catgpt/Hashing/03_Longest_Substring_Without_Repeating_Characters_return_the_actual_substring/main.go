package main

import "fmt"

func longestSubstringWithoutRepeating(s string) string {
	charindex := make(map[byte]int)
	maxLen := 0
	left := 0
	start := 0 // start index of the longest substring

	for right := 0; right < len(s); right++ {
		if idx, found := charindex[byte(right)]; found && idx >= left {
			left = idx + 1
		}
		charindex[s[right]] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}
	}
	return s[start : start+maxLen]
}

func main() {
	s := "abcabcbb"
	fmt.Println("Longest substring without repeating characters:", longestSubstringWithoutRepeating(s))
}

// ✅ This still runs in O(n) time and gives you the actual substring.
// Want to try the Unicode/rune version too (for full UTF-8 support like emojis 
// or non-English characters)? Or ready for another hashing challenge?