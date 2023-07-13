package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}
	fmt.Println(longestPalindrome(words))
}

func longestPalindrome(words []string) int {
	var newWords []string
	var palindromes []string
	length := len(words)
	for length > 1 {
		length = len(words)
		word := words[0]
		for i := 1; i < length; i++ {
			if word == reverse(words[i]) {
				newWords = insert(newWords, word, words[i])
				words = removeFromSlice(words, i)
				words = removeFromSlice(words, 0)
				break
			}
		}
		words = removeFromSlice(words, 0)
		if word == reverse(word) {
			palindromes = append(palindromes, word)
		}
	}
	if words[0] == reverse(words[0]) {
		palindromes = append(palindromes, words[0])
	}
	newWords = insert(newWords, longestWord(palindromes))
	return len(strings.Join(newWords, ""))
}

func longestWord(words []string) string {
	var word string
	for _, e := range words {
		if len(word) < len(e) {
			word = e
		}
	}
	return word
}

func insert(words []string, val ...string) []string {
	if words == nil {
		words = append([]string{}, val...)
	} else {
		ref := len(words) / 2
		head, tail := append([]string{}, words[0:ref]...), append([]string{}, words[ref:]...)
		words = append(head, val...)
		words = append(words, tail...)
	}
	return words
}

func reverse(s string) string {
	var result string
	length := len(s) - 1
	for i := length; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func removeFromSlice(words []string, i int) []string {
	var tmp []string
	tmp = append(tmp, words[0:i]...)
	tmp = append(tmp, words[i+1:]...)
	return tmp
}
