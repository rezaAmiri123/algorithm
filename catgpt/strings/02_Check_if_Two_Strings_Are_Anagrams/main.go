package main

import (
	"fmt"
	"strings"
)

// 🔸 Next Challenge: Check if Two Strings Are Anagrams
// Write a function that checks whether two strings are anagrams
// (case-insensitive, ignore spaces).

// Example:
// Input: "listen", "silent" → true
// Input: "Hello", "Olelh" → true
// Input: "test", "ttew" → false

// 🔍 Explanation:
// Normalize strings to lowercase and remove spaces.
// Use a hash map to count how many times each character appears in s1,
// then subtract counts using s2.
// If at any point a count goes negative or characters don’t match up, it’s not an anagram.

func IsAnagram(s1, s2 string) bool {
	// Normalize: lower case and remove spaces
	s1 = strings.ReplaceAll(strings.ToLower(s1)," ","")
	s2 = strings.ReplaceAll(strings.ToLower(s2)," ","")
	
	// If lengths differ, they can't be anagrams
	if len(s1)!= len(s2){
		return false
	}

	// Use a frequency map
	freq := make(map[rune]int)

	for _, ch := range s1{
		freq[ch]++
	}
	for _,ch := range s2{
		freq[ch]--
		if freq[ch]<0{
			return false
		}
	}
	return true
}

func main(){
	fmt.Println("listen","silent",IsAnagram("listen", "silent"))
	fmt.Println("Hello","Olelh",IsAnagram("Hello","Olelh"))
	fmt.Println("test","ttew",IsAnagram("test","ttew"))
}




