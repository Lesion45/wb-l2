package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSorter(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		options  SortOptions
		expected []string
	}{
		{
			name:     "SortNumeric",
			lines:    []string{"10 apple", "2 banana", "1 cherry"},
			options:  SortOptions{Column: 1, Numeric: true},
			expected: []string{"1 cherry", "2 banana", "10 apple"},
		},
		{
			name:     "SortReverse",
			lines:    []string{"apple", "banana", "cherry"},
			options:  SortOptions{Reverse: true},
			expected: []string{"cherry", "banana", "apple"},
		},
		{
			name:     "UniqueLines",
			lines:    []string{"apple", "banana", "apple", "cherry"},
			options:  SortOptions{Unique: true},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "SortByMonth",
			lines:    []string{"banana ", " apple", "cherry"},
			options:  SortOptions{IgnoreTrailing: true},
			expected: []string{" apple", "banana ", "cherry"},
		}, {
			name:     "SortIgnoreTrailingSpaces",
			lines:    []string{"banana ", " apple", "cherry"},
			options:  SortOptions{IgnoreTrailing: true},
			expected: []string{" apple", "banana ", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := NewSorter(tt.lines, tt.options)
			actual := sorter.Sort()

			assert.Equal(t, tt.expected, actual)
		})
	}
}
