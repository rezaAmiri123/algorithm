package main

import "strings"

// Sparse Search: Given a sorted array of strings that is interspersed with empty strings, write a
// method to find the location of a given string.
// EXAMPLE
// Input: ball, {"at", "" , "" , "" , "ball", ,,,, , ,,,, , "car", "" , "" , "dad", "" , ""}
// Output: 4

// SOLUTION
// If it weren't for the empty strings, we could simply use binary search. We would compare the string to be
// found, s tr, with the midpoint of the array, and go from there.
// With empty strings interspersed, we can implement a simple modification of binary search. All we need to
// do is fix the comparison against mid, in case mid is an empty string. We simply move mid to the closest
// non-empty string.
// The recursive code below to solve this problem can easily be modified to be iterative. We provide such an
// implementation in the code attachment.
func search(strs []string, str string, first, last int) int {
	if first > last {
		return -1
	}
	// Move mid to the middle
	mid := (last + first) / 2

	// If mid is empty, find closest non-empty string.
	if strs[mid] == "" {
		left := mid - 1
		right := mid + 1
		for {
			if left < first && right > last {
				return -1
			} else if right <= last && strs[right] != "" {
				mid = right
				break
			} else if left >= first && strs[left] != "" {
				mid = left
				break
			}
			right++
			left--
		}
	}
	// If mid is empty, find closest non-empty string.
	if str == strs[mid] { // Found it!
		return mid
	} else if strings.Compare(strs[mid], str) < 0 { // search right
		return search(strs, str, mid+1, last)
	} else { // search left
		return search(strs, str, first, mid-1)
	}
}

func Search(strs []string, str string) int {
	if strs == nil || str == "" {
		return -1
	}
	return search(strs, str, 0, len(strs)-1)
}

// The worst-case runtime for this algorithm is O( n). In fact, it's impossible to have an algorithm for this
// problem that is better than O(n) in the worst case. After all, you could have an array of all empty strings
// except for one non-empty string. There is no "smart" way to find this non-empty string. In the worst case,
// you will need to look at every element in the array.
// Careful consideration should be given to the situation when someone searches for the empty string. Should
// we find the location (which is an O( n) operation)? Or should we handle this as an error?

// There's no correct answer here. This is an issue you should raise with your interviewer. Simply asking this
// question will demonstrate that you are a careful coder.

// 413