package main

import "fmt"

// 🔄 Goal:
// Return all sequences tied for the longest consecutive streak — not just the first one.

// 🧱 Plan Update:
// Instead of tracking a single longest sequence:
// Track the max length.
// Track a slice of sequences ([][]int) that match the max length.
// Every time we build a sequence:
// If it's longer than current max → reset and add it.
// If it's equal to max → add it to the list.

func allLongestConsecutiveSequences(nums []int) [][]int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	var results [][]int
	maxLen := 0
	for num := range numSet {
		// Only start building sequence if num-1 doesn't exist
		if !numSet[num-1] {
			currentNum := num
			currentSeq := []int{currentNum}

			for numSet[currentNum+1] {
				currentNum++
				currentSeq = append(currentSeq, currentNum)
			}

			if len(currentSeq) > maxLen {
				maxLen = len(currentSeq)
				results = [][]int{currentSeq} //reset
			} else if len(currentSeq) == maxLen {
				results = append(results, currentSeq) // add another
			}
		}
	}
	return results
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2, 10, 11, 12, 13}
	sequences := allLongestConsecutiveSequences(nums)
	fmt.Println("All longest consecutive sequences:")
	for _, seq := range sequences {
		fmt.Println(seq)
	}
}

// Pretty slick, right? 😎
// Ready to try Top K Frequent Elements next, 
// or want to twist this with sorted output or return starting/ending values only?
