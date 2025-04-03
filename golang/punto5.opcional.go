package main

import (
	"fmt"
	"sort"
	"strings"
)

func isPalindrome(word string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	runeword := []rune(cleaned)
	for i, j := 0, len(runeword)-1; i < j; i, j = i+1, j-1 {
		if runeword[i] != runeword[j] {
			return false
		}
	}
	return true
}

func isAnagram(word1, word2 string) bool {
	normalize := func(word string) string {
		cleaned := strings.ToLower(strings.ReplaceAll(word, " ", ""))
		runes := []rune(cleaned)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		return string(runes)
	}
	return normalize(word1) == normalize(word2)
}

func isIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, char := range strings.ToLower(word) {
		if letters[char] {
			return false
		}
		letters[char] = true
	}
	return true
}

func main() {
	// Dificultad extra
	fmt.Println(isPalindrome("radar"))         // true
	fmt.Println(isAnagram("listen", "silent")) // true
	fmt.Println(isIsogram("background"))       // true
}
