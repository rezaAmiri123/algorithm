package main

import "fmt"

// 🔹 Question: Valid Parentheses
// Given a string s containing just the characters '(', ')', '{', '}', '[', and ']',
// determine if the input string is valid.
// A string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

func isValid(s string) bool {
	stack := []rune{}
	breacketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != breacketMap[ch] {
				return false
			}
			stack = stack[:len(stack)-1] //pop
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}

// 🧠 Explanation
// We use a stack to store opening brackets as we see them.
// When we hit a closing bracket, we check if it matches the last opened bracket 
// (top of the stack).
// If it doesn't match or the stack is empty when it shouldn't be — it's invalid.
// In the end, the stack must be empty for the string to be valid.


// 💡 Follow-Up Interview Ideas:
// Implement a Min Stack (stack that retrieves minimum in O(1)).
// Evaluate a Reverse Polish Notation expression.
// Simulate Undo/Redo operations.
// Decode a nested string pattern like "3[a2[c]]" → "accaccacc"
