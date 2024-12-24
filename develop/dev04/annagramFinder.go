package main

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func findAnagrams(words *[]string) *map[string][]string {
	anagrams := make(map[string][]string)
	wordMap := make(map[string]string)
	wordSet := make(map[string]struct{})

	for _, word := range *words {
		lowerWord := strings.ToLower(word)
		sortedWord := sortString(lowerWord)

		if _, exists := wordSet[lowerWord]; exists {
			continue
		}
		wordSet[lowerWord] = struct{}{}
		if originalWord, exists := wordMap[sortedWord]; exists {
			anagrams[originalWord] = append(anagrams[originalWord], lowerWord)
		} else {
			wordMap[sortedWord] = lowerWord
			anagrams[lowerWord] = []string{lowerWord}
		}
	}

	for key, value := range anagrams {
		if len(value) < 2 {
			delete(anagrams, key)
		} else {
			sort.Strings(value)
		}
	}

	return &anagrams
}
