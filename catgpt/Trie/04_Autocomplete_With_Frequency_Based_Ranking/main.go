package main

import (
	"fmt"
	"sort"
)

// 🔸 Objective: Autocomplete With Frequency-Based Ranking
// Each inserted word has a frequency (e.g., how often it's been typed/searched).
// Your autocomplete function should return results ordered by highest frequency first
// (and alphabetically for ties).

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	word     string
	freq     int
}

// Trie structure
type Trie struct {
	root *TrieNode
}

// Constructor
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert adds a word with a given frequency
func (t *Trie) Insert(word string, frequency int) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.word = word
	node.freq = frequency
}

// Autocomplete returns top words starting with prefix,
// sorted by frequency and alphabetically
func (t *Trie) Autocomplete(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return []string{}
		}
		node = node.children[ch]
	}
	var results []TrieNode
	t.dfs(node, &results)

	// Sort by frequency (desc), then alphabetically
	sort.Slice(results, func(i, j int) bool {
		if results[i].freq == results[j].freq {
			return results[i].word < results[j].word
		}
		return results[i].freq > results[j].freq
	})

	// Extract only the words
	words := make([]string, len(results))
	for i, n := range results {
		words[i] = n.word
	}

	return words
}

func (t *Trie) dfs(node *TrieNode, results *[]TrieNode) {
	if node.isEnd {
		*results = append(*results, *node)
	}
	for _, child := range node.children {
		t.dfs(child, results)
	}
}

// Test
func main() {
	trie := NewTrie()
	trie.Insert("app", 10)
	trie.Insert("apple", 5)
	trie.Insert("apply", 3)
	trie.Insert("apt", 8)
	trie.Insert("april", 2)

	fmt.Println(trie.Autocomplete("ap"))  // [app apt apple apply april]
	fmt.Println(trie.Autocomplete("app")) // [app apple apply]
}

// 🧠 Explanation
// We store freq on each node when the word ends.
// During autocomplete, we collect all matching nodes, then sort them:
// First by descending frequency.
// Then alphabetically.

