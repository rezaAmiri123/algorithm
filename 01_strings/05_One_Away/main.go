package main

import (
	"fmt"
	"math"
)

// One Away: There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.
// EXAMPLE
// pale, ple -> true
// pales, pale -> true
// pale,bale -> true
// pale,bae -> false

// SOLUTION
// There is a "brute force" algorithm to do this. We could check all possible strings that are one edit away by
// testing the removal of each character (and comparing), testing the replacement of each character (and
// comparing), and then testing the insertion of each possible character (and comparing).
// That would be too slow, so let's not bother with implementing it.
// This is one of those problems where it's helpful to think about the "meaning" of each of these operations.
// What does it mean for two strings to be one insertion, replacement, or removal away from each other?
// Replacement: Consider two strings, such as bale and pale, that are one replacement away. Yes, that
// does mean that you could replace a character in bale to make pale. But more precisely, it means that
// they are different only in one place.
// •
// Insertion: The strings apple and aple are one insertion away. This means that if you compared the
// strings, they would be identical-except for a shift at some point in the strings.
// •
// Removal: The strings apple and aple are also one removal away, since removal is just the inverse of
// insertion.
// We can go ahead and implement this algorithm now. We'll merge the insertion and removal check into one
// step, and check the replacement step separately.
// Observe that you don't need to check the strings for insertion, removal, and replacement edits. The lengths
// of the strings will indicate which of these you need to check.
func oneEditAway(first, second string) bool {
	if len(first) == len(second) {
		return oneEditReplace(first, second)
	} else if len(first)+1 == len(second) {
		return oneEditInsert(first, second)
	} else if len(first)-1 == len(second) {
		return oneEditInsert(second, first)
	}
	return false
}

func oneEditReplace(first, second string) bool {
	foundDiffrence := false
	for i, _ := range first {
		if first[i] != second[i] {
			if foundDiffrence {
				return false
			}
			foundDiffrence = true
		}
	}
	return true
}

/* Check if you can insert a character into sl to make s2. */
func oneEditInsert(str1, str2 string) bool {
	index1 := 0
	index2 := 0
	for index2 < len(str2) && index1 < len(str1) {
		if str1[index1] != str2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}

// This algorithm (and almost any reasonable algorithm) takes O ( n) time, where n is the length of the shorter
// string.
// Why is the runtime dictated by the shorter string instead of the longer string? If the strings are
// the same length (plus or minus one character), then it doesn't matter whether we use the longer
// string or the shorter string to define the runtime. If the strings are very different lengths, then the
// algorithm will terminate in 0( 1) time. One really, really long string therefore won't significantly
// extend the runtime. It increases the runtime only if both strings are long.

// We might notice that the code for one Edi tReplace is very simtlar to that for one Edi tinsert. We can
// merge them into one method.
// To do this, observe that both methods follow similar logic: compare each character and ensure that the
// strings are only different by one. The methods vary in how they handle that difference. The method
// one Edi tReplace does nothing other than flag the difference, whereas one Edi tinsert increments
// the pointer to the longer string. We can handle both of these in the same method.

func oneEditAwayImproved(first, second string) bool {
	// Length checks.
	if math.Abs(float64(len(first))-float64(len(second))) > 1 {
		return false
	}

	// Get shorter and longer string.
	var str1, str2 string
	if len(first) < len(second) {
		str1 = first
		str2 = second
	} else {
		str1 = second
		str2 = first
	}

	index1 := 0
	index2 := 0
	foundDiffrence := false
	for index2 < len(str2) && index1 < len(str1) {
		if str1[index1] != str2[index2] {
			// Ensure that this is the first difference found
			if foundDiffrence {
				return false
			}
			foundDiffrence = true

			// On replace, move shorter pointer
			if len(str1) == len(str2) {
				index1++
			}
		} else {
			index1++ // If matching, move shorter pointer
		}
		index2++ // Always move pointer for longer string
	}
	return true
}

func main() {
	fmt.Println(oneEditAway("pale", "bale"))
	fmt.Println(oneEditAway("pale", "bae"))
	fmt.Println("improved")
	fmt.Println(oneEditAwayImproved("pale", "bale"))
	fmt.Println(oneEditAwayImproved("pale", "bae"))
}

// Some people might argue the first approach is better, as it is clearer and easier to follow. Others, however,
// will argue that the second approach is better, since it's more compact and doesn't duplicate code (which
// can facilitate maintainability).
// You don't necessarily need to "pick a side:'You can discuss the tradeoffs with your interviewer.
