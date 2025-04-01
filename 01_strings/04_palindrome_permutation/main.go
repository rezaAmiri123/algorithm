package main

import (
	"fmt"
	"strings"
)

// A palindrome is a word or phrase that is the same forward and backward

// Solution 1: we eed to have an even number of almost all characters
// and at most one character(the middle character) can have an odd number
// What does it take to be able to write a set of characters the same way forwards and backwards? We need to
// have an even number of almost all characters, so that half can be on one side and half can be on the other
// side. At most one character (the middle character) can have an odd count.
// For example, we know tactcoapapa is a permutation of a palindrome because it has two Ts, four As, two
// Cs, two Ps, and one 0. That O would be the center of all possible palindromes.

// To be more precise, strings with even length (after removing all non-letter characters) must have
// all even counts of characters. Strings of an odd length must have exactly one character with
// an odd count. Of course, an "even" string can't have an odd number of exactly one character,
// otherwise it wouldn't be an even-length string (an odd number+ many even numbers= an odd
// number). Likewise, a string with odd length can't have all characters with even counts (sum of
// evens is even). It's therefore sufficient to say that, to be a permutation ot a palindrome, a string
// can have no more than one character that is odd. This will cover both the odd and the even cases.

// Solution#1
// Implementing this algorithm is fairly straightforward. We use a hash table to count how many times each
// character appears. Then, we iterate through the hash table and ensure that no more than one character has
// an odd count.

func isPermutationOfPalindrome(str string) bool {
	a := rune('a')
	z := rune('z')
	charMap := make(map[rune]int)
	for _, char := range str {
		charNumber := rune(char)
		if charNumber >= a && charNumber <= z {
			charMap[char]++
		}
	}
	foundOdd := false
	for _, value := range charMap {
		if value%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return foundOdd
}

// Solution #2
// We can't optimize the big O time here since any algorithm will always have to look through the entire
// string. However, we can make some smaller incremental improvements. Because this is a relatively simple
// problem, it can be worthwhile to discuss some small optimizations or at least some tweaks.
// Instead of checking the number of odd counts at the end, we can check as we go along. Then, as soon as
// we get to the end, we have our answer.
func isPermutationOfPalindromeImproved(str string) bool {
	oddCount := 0
	a := rune('a')
	z := rune('z')
	charMap := make(map[rune]int)
	for _, char := range str {
		charNumber := rune(char)
		if charNumber >= a && charNumber <= z {
			charMap[char]++
			if charMap[char]%2 == 1 {
				oddCount++
			} else {
				oddCount--
			}
		}
	}

	return oddCount <= 1
}

func main() {
	// fmt.Println(isPermutationOfPalindrome(strings.ToLower("123rfdfdg")))
	fmt.Println(isPermutationOfPalindrome(strings.ToLower("Tact Coa")))
	fmt.Println(isPermutationOfPalindrome(strings.ToLower("Tact Coaa")))
	fmt.Println("improved")
	fmt.Println(isPermutationOfPalindromeImproved(strings.ToLower("Tact Coa")))
	fmt.Println(isPermutationOfPalindromeImproved(strings.ToLower("Tact Coaa")))

	// also there is another sulotion at page 209
}
