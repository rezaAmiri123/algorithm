package main

import "fmt"

// string 05
func CountPalindromicSubstrings(s string)int{
	count :=0
	n := len(s)

	for center := 0;center<n*2-1;center++{
		left :=center/2
		right := left +center%2

		for left>=0&&right<n&&s[left]==s[right]{
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
