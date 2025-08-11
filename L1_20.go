package main

import "fmt"

func reverseRunes(r []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	reverseRunes(runes, 0, len(runes)-1)
	startOfWord := 0
	for i := 0; i < len(runes); i++ {

		if runes[i] == ' ' {
			reverseRunes(runes, startOfWord, i-1)
			startOfWord = i + 1
		}
	}
	reverseRunes(runes, startOfWord, len(runes)-1)

	return string(runes)
}

func main() {
	input := "snow dog sun"
	output := reverseWords(input)
	fmt.Printf("'%s' -> '%s'\n", input, output)

}