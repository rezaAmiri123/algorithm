package main

import (
	"fmt"
	"sort"
)

// 🔑 Problem: Group Anagrams
// Given an array of strings, group the anagrams together.
// Anagram = same letters in a different order (e.g., "eat", "tea", "ate")

// Input:  strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// Output: [["eat","tea","ate"],["tan","nat"],["bat"]]

// 🧠 Hashing Strategy:
// We’ll use a map[string][]string, where:
// The key is a signature for the anagram group.
// The value is a list of words that belong to that anagram group.
// Two common ways to build a signature:
// Sorted string (easier and common)
// Character count array as string key (faster for large datasets)
// We'll use the sorted string method here.

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	grouped := groupAnagrams(strs)
	for _, group := range grouped {
		fmt.Println(group)
	}
}

// Let me know if you want:
// The character count key version (better performance),
// A twist like returning group sizes, or
// Another type of hashing challenge like Top K Frequent Elements 
// or Longest Consecutive Sequence.
