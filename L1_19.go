package main

import "fmt"
func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	str1 := "главрыба"
	reversed1 := reverse(str1)
	fmt.Printf("'%s' -> '%s'\n", str1, reversed1)

	str2 := "a, б, 貓, 🍌"
	reversed2 := reverse(str2)
	fmt.Printf("'%s' -> '%s'\n", str2, reversed2)
}