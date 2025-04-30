package main

import "fmt"

// 🔸 Count All Palindromic Substrings
// Problem:
// Given a string s, return the number of palindromic substrings in it.
// A string is a palindrome if it reads the same forward and backward.
// Every single character is a palindrome by default.


// Here’s the ✅ solution to Count All Palindromic Substrings in Go — 
// using the expand around center technique (which is both elegant and efficient):

func CountPalindromicSubstrings(s string)int{
	count := 0
	n := len(s)

	for center :=0;center<2*n-1;center++{
		left := center/2
		right := left+center%2

		for left >= 0 && right<n && s[left]==s[right]{
			count++
			left--
			right++
		}
	}
	return count
}


func main() {
	fmt.Println(CountPalindromicSubstrings("abc"))  // Output: 3 → "a", "b", "c"
	fmt.Println(CountPalindromicSubstrings("aaa"))  // Output: 6 → "a", "a", "a", "aa", "aa", "aaa"
	fmt.Println(CountPalindromicSubstrings("abba")) // Output: 6 → "a", "b", "b", "a", "bb", "abba"
}

// 🔍 Explanation:
// We treat every character and every gap between characters as a potential center of a palindrome:
// center = 0 to 2*n - 2
// For each center, expand outward while the characters on both sides match.
// Count each valid expansion as a palindromic substring.

// Why this works:
// Odd-length palindromes: "aba" (center at 'b')
// Even-length palindromes: "aa" (center between 'a' and 'a')

