package main

import (
	"fmt"
	"strings"
)

type FilterOptions struct {
	After       int
	Before      int
	Context     int
	Count       bool
	IgnoreCase  bool
	InvertMatch bool
	Fixed       bool
	LineNum     bool
	Pattern     string
}

type FilterResult struct {
	LineNum int
	Line    string
	Matches bool
}

func Filter(lines []string, options FilterOptions) []FilterResult {
	var results []FilterResult
	pattern := options.Pattern

	if options.IgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	for i, line := range lines {
		lineToMatch := line
		if options.IgnoreCase {
			lineToMatch = strings.ToLower(line)
		}

		matches := false
		if options.Fixed {
			matches = lineToMatch == pattern
		} else {
			matches = strings.Contains(lineToMatch, pattern)
		}

		if options.InvertMatch {
			matches = !matches
		}

		if matches {
			start := max(0, i-options.Before)
			end := min(len(lines), i+options.After+1)

			for j := start; j < end; j++ {
				results = append(results, FilterResult{
					LineNum: j + 1,
					Line:    lines[j],
					Matches: j == i,
				})
			}
		}
	}

	if options.Count {
		results = []FilterResult{{Line: fmt.Sprintf("%d", len(results))}}
	}

	return results
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
