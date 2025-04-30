package main

import "fmt"

// 🔸 Question: Implement a Word Dictionary with Wildcard Search
// Design a data structure that supports adding new words and finding 
// if a string matches any previously added word. 
// The string may contain the dot character '.', which can represent any one letter.
// Implement the following methods:
// AddWord(word string)
// Search(word string) bool — '.' can match any letter.

// TrieNode represents each node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// WordDictionary is the main structure
type WordDictionary struct {
	root *TrieNode
}

// Constructor
func NewWordDictionary() *WordDictionary {
	return &WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// AddWord inserts a word into the dictionary
func (wd *WordDictionary) AddWord(word string) {
	node := wd.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search returns true if the word is in the dictionary
func (wd *WordDictionary) Search(word string) bool {
	return searchHelper(word, 0, wd.root)
}

func searchHelper(word string, index int, node *TrieNode) bool {
	if index == len(word) {
		return node.isEnd
	}
	ch := rune(word[index])
	if ch == '.' {
		for _, child := range node.children {
			if searchHelper(word, index+1, child) {
				return true
			}
		}
		return false
	}

	child, exists := node.children[ch]
	if !exists {
		return false
	}
	return searchHelper(word, index+1, child)
}

// Test example
func main() {
	dict := NewWordDictionary()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")

	fmt.Println(dict.Search("pad")) // false
	fmt.Println(dict.Search("bad")) // true
	fmt.Println(dict.Search(".ad")) // true
	fmt.Println(dict.Search("b..")) // true
}

// 🧠 Explanation
// The dot '.' acts like a wildcard, 
// so the recursive searchHelper explores all children at that level.
// We only return true if one of those wildcard paths ends successfully.