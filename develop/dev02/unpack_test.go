package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input         string
		expected      string
		hasError      bool
		expectedError error
	}{
		{"a4bc2d5e", "aaaabccddddde", false, nil},
		{"abcd", "abcd", false, nil},
		{"45", "", true, ErrInvalidString},
		{"", "", false, nil},
		{"qwe\\4\\5", "qwe45", false, nil},
		{"qwe\\45", "qwe44444", false, nil},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false, nil},
		{"\\", "", true, ErrWrongEscapeSeq},
		{"a\\", "", true, ErrWrongEscapeSeq},
		{"a3\\", "", true, ErrWrongEscapeSeq},
	}

	for _, test := range tests {
		result, err := Unpack(test.input)

		if test.hasError {
			assert.Error(t, err, "Input: %s", test.input)
			assert.ErrorIs(t, err, test.expectedError, "Input: %s", test.input)
		} else {
			assert.NoError(t, err, "Input: %s", test.input)
			assert.Equal(t, test.expected, result, "Input: %s", test.input)
		}
	}
}
