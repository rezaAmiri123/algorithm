package main

import "fmt"

// 🔹 Reverse a String (without using built-in reverse)
// Write a function that reverses a string in Go.

// Example:
// Input: "hello"
// Output: "olleh"

// Let me know if you want a hint, solution, or you're ready to try it.
// And once you're done, I’ll throw you a slightly harder one like checking for anagram,
// longest substring without repeating characters, etc. You in?

func ReverseString(s string)string{
	// Convert string to rune slice to handle Unicode properly
	runes := []rune(s)
	n := len(runes)

	// Swap characters from both ends
	for i:=0;i<n/2;i++{
		runes[i], runes[n-1-i]=runes[n-1-i],runes[i]
	}

	return string(runes)
}

// 🔍 Explanation:
// Go strings are UTF-8 encoded, so using []rune helps properly handle characters like emojis 
// or non-ASCII characters.

// We swap runes in place from both ends toward the center.
func main(){
	fmt.Println(ReverseString("hello"))
}