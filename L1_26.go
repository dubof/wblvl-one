package main

import (
	"fmt"
	"strings"
)

func areCharsUnique(s string) bool {
	lowerStr := strings.ToLower(s)
	seen := make(map[rune]struct{})
	for _, r := range lowerStr {
		if _, ok := seen[r]; ok {
			return false
		}
		seen[r] = struct{}{}
	}
	return true
}

func main() {
	fmt.Printf("'%s' -> %t\n", "abcd", areCharsUnique("abcd"))
	fmt.Printf("'%s' -> %t\n", "abCdefAaf", areCharsUnique("abCdefAaf"))
	fmt.Printf("'%s' -> %t\n", "aabcd", areCharsUnique("aabcd"))
	fmt.Printf("'%s' -> %t\n", "abcde", areCharsUnique("abcde"))
	fmt.Printf("'%s' -> %t\n", "abCde", areCharsUnique("abCde"))
}