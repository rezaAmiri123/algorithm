package main

import "fmt"

// Longest Substring Without Repeating Characters
// Problem:
// Given a string, find the length of the longest substring without repeating characters.

// 🔍 Explanation:
// charIndex stores the last seen index of each character.
// If a character repeats within the current window,
// move the start pointer past the last occurrence.
// Keep updating the maxLength as you go.

// Example:
// Input: "abcabcbb"
// → Output: 3 ("abc")

// Input: "bbbbb"
// → Output: 1 ("b")

// Input: "pwwkew"
// → Output: 3 ("wke")
func LengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int)
    start, maxLength := 0, 0
    runes := []rune(s)

    for i, ch := range runes {
        if lastIdx, found := charIndex[ch]; found && lastIdx >= start {
            start = lastIdx + 1
        }
        charIndex[ch] = i
        if i-start+1 > maxLength {
            maxLength = i - start + 1
        }
    }

    return maxLength
}

func main(){
	fmt.Println("abcabcbb", LengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbb", LengthOfLongestSubstring("bbbb"))
	fmt.Println("pwwkew", LengthOfLongestSubstring("pwwkew"))
}