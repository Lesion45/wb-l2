package dev06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name           string
		line           string
		options        CutOptions
		expectedResult string
	}{
		{
			name:           "Split by TAB and select fields 1,2",
			line:           "apple\torange\tbanana",
			options:        CutOptions{Fields: "1,2", Delimiter: "\t"},
			expectedResult: "apple\torange",
		},
		{
			name:           "Split by comma and select fields 1,3",
			line:           "apple,orange,banana",
			options:        CutOptions{Fields: "1,3", Delimiter: ","},
			expectedResult: "apple,banana",
		},
		{
			name:           "Only fields 2 and 3 with comma delimiter",
			line:           "cat,dog,fish",
			options:        CutOptions{Fields: "2,3", Delimiter: ","},
			expectedResult: "dog,fish",
		},
		{
			name:           "Only lines with delimiter (ignore non-delimited)",
			line:           "hello world",
			options:        CutOptions{Fields: "1", Delimiter: ",", Separated: true},
			expectedResult: "",
		},
		{
			name:           "Valid line with tab separator",
			line:           "apple\tbanana",
			options:        CutOptions{Fields: "1", Delimiter: "\t"},
			expectedResult: "apple",
		},
		{
			name:           "Line without selected field index",
			line:           "apple,banana",
			options:        CutOptions{Fields: "3", Delimiter: ","},
			expectedResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := processLine(tt.line, &tt.options)
			assert.Equal(t, tt.expectedResult, actual)
		})
	}
}
