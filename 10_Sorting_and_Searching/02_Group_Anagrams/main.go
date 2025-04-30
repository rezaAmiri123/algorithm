package main

// Group Anagrams: Write a method to sort an array ot strings so that all tne anagrnms are next to
// each other.

// SOLUTION
// This problem asks us to group the strings in an array such that the anagrams appear next to each other.
// Note that no specific ordering of the words is required, other than this.
// We need a quick and easy way of determining if two strings are anagrams of each other. What defines if two
// words are anagrams of each other? Well, anagrams are words that have the same characters but in different
// orders. It follows then that if we can put the characters in the same order, we can easily check if the new
// words are identical.
// One way to do this is to just apply any standard sorting algorithm, like merge sort or quick sort, and modify
// the comparator. This comparator will be used to indicate that two strings which are anagrams of each other
// are equivalent.

// What's the easiest way of checking if two words are anagrams? We could count the occurrences of the
// distinct characters in each string and return true if they match. Or, we could just sort the string. After all,
// two words which are anagrams will look the same once they're sorted.
// The code below implements the comparator.
type AnagramComparator struct{

}

// to be continued at page 408
