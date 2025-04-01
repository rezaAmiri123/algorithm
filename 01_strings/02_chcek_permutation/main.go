package main

import (
	"fmt"
	"sort"
)

// Given two strings, write a method to decide if one is a permutation
// of the other

// questions:
// 1- is it ase sensitive --> yes
// 2- if whitespace is significcant--> "dog  " is different from "dog"

// solution1: sort the string
// if two strings are permutation, then we know they have the same characters, but in different orders
// so we sort them and compare the sorted versions of them
type sortString []rune

func (x sortString) Len() int           { return len(x) }
func (x sortString) Less(i, j int) bool { return x[i] < x[j] }
func (x sortString) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func permutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	sortedStr1 := []rune(str1)
	sortedStr2 := []rune(str2)
	sort.Sort(sortString(sortedStr1))
	sort.Sort(sortString(sortedStr2))
	if string(sortedStr1) != string(sortedStr2){
		return false
	}
	return true
}

// solution2: check if the two strings have idential character counts
// defination of permutation: two words with the same character counts  
func permutation2(str1, str2 string)bool{
	if len(str1) != len(str2) {
		return false
	}
	letters := make(map[rune]int)
	for _, char := range str1{
		letters[char]++
	}
	for _, char := range str2{
		count, exists := letters[char]
		if !exists{
			return false
		}
		if count == 0{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("solution 1")
	fmt.Println(permutation("acdb", "abcd"))
	fmt.Println(permutation("acdb", "efgh"))
	fmt.Println("solution 2")
	fmt.Println(permutation2("acdb", "abcd"))
	fmt.Println(permutation2("acdb", "efgh"))

}
