package main

import "fmt"

// 🔸 Implement strStr() (Needle in a Haystack)
// Problem:
// Implement strStr(haystack string, needle string) int — return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Constraints:
// Do not use strings.Index or other built-in search methods.
// If needle is an empty string, return 0.

func StrStr(haystack string, needle string) int {
	if needle == ""{
		return 0
	}

	n,m := len(haystack),len(needle)
	if m>n{
		return -1
	}

	for i:=0;i<=n-m;i++{
		if haystack[i:i+m]==needle{
			return i
		}
	}
	return -1	
}

func main() {
	fmt.Println(StrStr("hello", "ll"))  // 2
	fmt.Println(StrStr("abcdef", "gh")) // -1
	fmt.Println(StrStr("aaa", "a"))     // 0
	fmt.Println(StrStr("abc", ""))      // 0
}

// 🔍 Explanation:
// We loop through haystack with a window of length equal to needle.
// At each step, check if the substring matches needle.
// If yes → return the index.
// If the loop ends without a match → return -1.
// If needle is empty → by definition, return 0.

// Want a challenge now that uses a bit more thinking — 
// like counting all palindromic substrings, or implementing wildcard pattern matching (? and *)? 
// Or something in the "hard" tier like minimum window substring?
