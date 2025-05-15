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
// We treat every character and every gap between characters 
// as a potential center of a palindrome:
// center = 0 to 2*n - 2
// For each center, expand outward while the characters on both sides match.
// Count each valid expansion as a palindromic substring.

// Why this works:
// Odd-length palindromes: "aba" (center at 'b')
// Even-length palindromes: "aa" (center between 'a' and 'a')

// This Go code defines a function CountPalindromicSubstrings that counts how many palindromic substrings exist in a given string s. A palindromic substring is a sequence of characters that reads the same forward and backward (e.g., "a", "aa", "aba").

// Code Breakdown
// func CountPalindromicSubstrings(s string) int {
// This function takes a string s and returns an integer — 
// the total count of palindromic substrings.
// 	count := 0
// 	n := len(s)
// Initializes the counter and gets the length of the string.

// Loop Through All Possible Centers
// 	for center := 0; center < 2*n-1; center++ {
// This loop simulates all possible centers of substrings (both odd and even lengths). 
// Why 2*n - 1?

// A string of length n can have n odd-length centers (each character) 
// and n-1 even-length centers (between characters).

// Total centers = n + (n - 1) = 2n - 1.

// Expand Around Center

// 		left := center / 2
// 		right := left + center % 2
// For odd-length palindromes (like "aba"), left == right.

// For even-length palindromes (like "aa"), right == left + 1.

// 		for left >= 0 && right < n && s[left] == s[right] {
// 			count++
// 			left--
// 			right++
// 		}
// This inner loop expands outward from the center as long as the substring is still a palindrome.

// Each time s[left] == s[right], it’s a valid palindrome → increment count and keep expanding.

// Return the Result

// 	return count
// }
// Usage in main
// func main() {
// 	fmt.Println(CountPalindromicSubstrings("abc"))  // 3 → "a", "b", "c"
// 	fmt.Println(CountPalindromicSubstrings("aaa"))  // 6 → "a", "a", "a", "aa", "aa", "aaa"
// 	fmt.Println(CountPalindromicSubstrings("abba")) // 6 → "a", "b", "b", "a", "bb", "abba"
// }
// This demonstrates the function with example inputs.

// Summary
// Time complexity: O(n²), where n is the length of the string — 
// each center could expand up to n characters.

// Space complexity: O(1), constant space — no extra storage.

// Would you like a visualization of how the center expansion works for one of the examples?
