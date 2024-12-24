package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name     string
		lines    []string
		options  FilterOptions
		expected []FilterResult
	}{
		{
			name:  "Exact match",
			lines: []string{"hello", "world", "hello world"},
			options: FilterOptions{
				Pattern: "hello",
				Fixed:   true,
			},
			expected: []FilterResult{
				{LineNum: 1, Line: "hello", Matches: true},
			},
		},
		{
			name:  "Case insensitive match",
			lines: []string{"Hello", "world", "hello world"},
			options: FilterOptions{
				Pattern:    "hello",
				IgnoreCase: true,
			},
			expected: []FilterResult{
				{LineNum: 1, Line: "Hello", Matches: true},
				{LineNum: 3, Line: "hello world", Matches: true},
			},
		},
		{
			name:  "Invert match",
			lines: []string{"hello", "world", "hello world"},
			options: FilterOptions{
				Pattern:     "hello",
				InvertMatch: true,
			},
			expected: []FilterResult{
				{LineNum: 2, Line: "world", Matches: true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results := Filter(tc.lines, tc.options)
			assert.Equal(t, tc.expected, results)
		})
	}
}
