package main

import "fmt"

// 🧑‍💻 Goal:
// Instead of just returning the length of the longest consecutive sequence,
// we’ll return the actual sequence as a slice of integers.

// 🔁 Plan:
// Same as before, but with a twist:
// When we find the start of a sequence (num - 1 not in the set), we build the sequence into a slice.
// Track the sequence if its length is the longest so far.

func longestConsecutiveSequence(nums []int) []int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestSeq := []int{}

	for num := range numSet {
		// Check if it's the start of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentSeq := []int{currentNum}

			for numSet[currentNum+1] {
				currentNum++
				currentSeq = append(currentSeq, currentNum)
			}

			if len(currentSeq) > len(longestSeq) {
				longestSeq = currentSeq
			}
		}
	}
	return longestSeq
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Longest consecutive sequence:", longestConsecutiveSequence(nums))
}

// 💡 Bonus tip: If multiple sequences are tied for max length, 
// this version will return the first one found.

// Want to add support for returning all sequences tied for longest, 
// or feeling like jumping to Top K Frequent Elements or Subarray Sum Equals K?