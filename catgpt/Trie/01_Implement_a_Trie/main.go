package main

import "fmt"

// 🔹 Question: Implement a Trie (Prefix Tree)
// Design and implement a trie with the following methods:
// Insert(word string): Inserts a word into the trie.
// Search(word string) bool: Returns true if the word is in the trie.
// StartsWith(prefix string) bool:
// Returns true if there is any word in the trie that starts with the given prefix.

// TrieNode represents each node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie is the main structure
type Trie struct {
	root *TrieNode
}

// Constructor
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert adds a word into the trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search returns true if the word is in the trie
func (t *Trie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.isEnd
}

// StartsWith returns true if any word starts with the prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

// Helper function to find the end node of a prefix/word
func (t *Trie) findNode(s string) *TrieNode {
	node := t.root
	for _, ch := range s {
		if _, exists := node.children[ch]; !exists {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

// Testing the Trie
func main() {
	trie := NewTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}

// 🧠 Explanation
// A TrieNode keeps track of:
// children: A map of characters to child TrieNodes.
// isEnd: A flag to mark if this node completes a valid word.
// Insert: Adds a word character-by-character, creating new nodes as needed.
// Search: Uses the helper findNode to get to the node that would be the end of the word and checks isEnd.
// StartsWith: Uses findNode but only checks if such a prefix path exists.
