package main

import "fmt"

// 🧠 Problem: Detect the Rightmost Set Bit
// Prompt:
// Given an integer n, return an integer that has only the rightmost set bit of n turned on
// (i.e., all other bits are 0).

// ✅ Example:
// Input:  n = 12  // binary: 1100
// Output: 4       // binary: 0100

// 🔍 Explanation:
// To isolate the rightmost set bit, you can use this trick:
// rightmostSetBit := n & -n

// This works because of two's complement representation.
// -n is the two's complement of n, which flips the bits and adds 1.
// n & -n leaves only the lowest bit that was set to 1.

func rightmostSetBit(n int) int {
	return n & -n
}

func main() {
	n := 12
	fmt.Printf("Rightmost set bit of %d is: %d\n", n, rightmostSetBit(n))
}

// Would you like to try a follow-up where we count the number of set bits 
// (also known as Hamming Weight)?