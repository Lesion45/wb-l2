package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "Multiple anagrams",
			words: []string{"listen", "silent", "enlist", "inlets", "stone"},
			expected: map[string][]string{
				"listen": {"enlist", "inlets", "listen", "silent"},
				"silent": {"enlist", "inlets", "listen", "silent"},
				"enlist": {"enlist", "inlets", "listen", "silent"},
				"inlets": {"enlist", "inlets", "listen", "silent"},
			},
		},
		{
			name:     "Unique words",
			words:    []string{"apple", "banana", "cherry", "date"},
			expected: map[string][]string{},
		},
		{
			name:     "Empty list",
			words:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "Single word",
			words:    []string{"apple"},
			expected: map[string][]string{},
		},
		{
			name:  "Case insensitive anagrams",
			words: []string{"Listen", "Silent", "inlets", "enlist"},
			expected: map[string][]string{
				"listen": {"enlist", "inlets", "listen", "silent"},
				"silent": {"enlist", "inlets", "listen", "silent"},
				"enlist": {"enlist", "inlets", "listen", "silent"},
				"inlets": {"enlist", "inlets", "listen", "silent"},
			},
		},
		{
			name:  "Anagrams with different lengths",
			words: []string{"abc", "bac", "abcd", "abdc", "cab", "dabc"},
			expected: map[string][]string{
				"abc":  {"abc", "bac", "cab"},
				"bac":  {"abc", "bac", "cab"},
				"cab":  {"abc", "bac", "cab"},
				"abcd": {"abcd", "abdc", "dabc"},
				"abdc": {"abcd", "abdc", "dabc"},
				"dabc": {"abcd", "abdc", "dabc"},
			},
		},
		{
			name:     "Repeated word",
			words:    []string{"apple", "apple", "apple", "apple"},
			expected: map[string][]string{},
		},
		{
			name:  "Duplicate anagrams",
			words: []string{"eat", "tea", "ate", "tea", "ate"},
			expected: map[string][]string{
				"ate": {"ate", "eat", "tea"},
				"eat": {"ate", "eat", "tea"},
				"tea": {"ate", "eat", "tea"},
			},
		},
		{
			name:     "Empty strings",
			words:    []string{"", "", " "},
			expected: map[string][]string{},
		},
		{
			name:  "Large words",
			words: []string{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
			expected: map[string][]string{
				"abcdefghijklmnopqrstuvwxyz": {"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
				"zyxwvutsrqponmlkjihgfedcba": {"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findAnagrams(&tt.words)
			assert.Equal(t, tt.expected, *actual)
		})
	}
}
