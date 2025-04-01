package main

import "fmt"

// Implement an algorithm to determine is a string has all unique characters
// what if you cannot use aditional data structures?

func isUniqueChars(str string)bool{
	charSet := make(map[rune]struct{})
	for _, char := range str{
		_, exists := charSet[char]
		if exists{
			return false
		}
		charSet[char] = struct{}{}
	}
	return true
}

func main(){
	str1 := "123edftg"
	str2 := "123edftggggggggg"
	fmt.Printf("is %s unique: %v\n", str1, isUniqueChars(str1))
	fmt.Printf("is %s unique: %v\n", str2, isUniqueChars(str2))
}