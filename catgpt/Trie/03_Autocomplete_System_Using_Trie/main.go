package main

import (
	"fmt"
	"sort"
)

// 🔹 Question: Autocomplete System Using Trie
// Implement a function that returns autocomplete suggestions based on a given prefix.
// You’ll:
// Insert a list of words into a Trie.
// Implement a method Autocomplete(prefix string) []string that
// returns all words in the Trie starting with the given prefix.

// TrieNode structure
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	word     string
}

// Trie structure
type Trie struct {
	root *TrieNode
}

// Constructor
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.word = word // store full word for easy lookup
}

// Autocomplete returns all words starting with a given prefix
func (t *Trie) Autocomplete(perfix string) []string {
	node := t.root
	for _, ch := range perfix {
		if _, exists := node.children[ch]; !exists {
			return []string{}
		}
		node = node.children[ch]
	}

	// DFS to find all words below this node
	var results []string
	t.dfs(node, &results)

	sort.Strings(results) // Optional: sort alphabetically
	return results
}

// DFS helper to collect words
func (t *Trie) dfs(node *TrieNode, results *[]string) {
	if node.isEnd {
		*results = append(*results, node.word)
	}
	for _, child := range node.children {
		t.dfs(child, results)
	}
}

// Test
func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "apply", "apt", "bat", "bath", "banana"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Autocomplete("ap"))  // [app apple apply apt]
	fmt.Println(trie.Autocomplete("ba"))  // [banana bat bath]
	fmt.Println(trie.Autocomplete("z"))   // []
	fmt.Println(trie.Autocomplete("app")) // [app apple apply]
}

// 🧠 Explanation
// We traverse the Trie until the end of the prefix.
// From that node, we DFS to collect all complete words (isEnd == true) below it.
// We store the whole word in the word field for easy access during DFS.

// 💡 Extensions You Could Be Asked:
// Limit results to top k suggestions.
// Add frequency scoring (most common words first).
// Return suggestions as-you-type in an app simulation.
