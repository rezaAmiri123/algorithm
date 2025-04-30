package main

import "strings"

// Word Frequencies: Design a method to find the frequency of occurrences of any given word in a
// book. What if we were running this algorithm multiple times?

// SOLUTION
// Let's start with the simple case.
// Solution: Single Query
// In this case, we simply go through the book, word by word, and count the number of times that a word
// appears. This will take O ( n) time. We know we can't do better than that since we must look at every word
// in the book.

func getFrequency(book []string, word string)int{
	word = strings.ToLower(strings.Trim(word," "))
	count := 0
	for _, w := range book{
		if strings.ToLower(strings.Trim(w," ")) == word{
			count++
		}
	}
	return count
}

// We have also converted the string to lowercase and trimmed it. You can discuss with your interviewer if this
// is necessary (or even desired).

// continued at page 474
