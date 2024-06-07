package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var word string
	fmt.Print("Enter word: ")
	fmt.Scan(&word)
	fmt.Println(IsIsogram(word))
}

func IsIsogram(word string) bool {
	mapChars := make(map[rune]int)
	word = strings.ToLower(word)
	for _, val := range word {
		if unicode.IsLetter(val) {
			mapChars[val]++
		}
	}
	for _, char := range mapChars {
		if char > 1 {
			return false
		}
	}
	return true
}
