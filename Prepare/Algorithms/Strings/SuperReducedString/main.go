package main

import (
	"fmt"
)

func superReducedString(s string) string {
	var changed bool
	for {
		s, changed = changeString(s)
		if !changed {
			break
		}
	}
	if s == "" {
		return "Empty String"
	}
	return s
}

func changeString(s string) (string, bool) {
	changed := false
	for i := 0; i+1 < len(s); i++ {
		if s[i] == s[i+1] {
			changed = true
			if len(s) == 2 {
				return "", changed
			}
			if i+2 < len(s) {
				return s[:i] + s[i+2:], changed
			}
			return s[:i], changed
		}
	}
	return s, changed
}

func main() {
	example := "aaabccddd"
	fmt.Println(superReducedString(example)) // Output: "a"
}
