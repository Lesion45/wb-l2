package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString  = errors.New("invalid string")
	ErrWrongEscapeSeq = errors.New("wrong escape sequence")
)

func Unpack(input string) (string, error) {
	var result strings.Builder
	var prevRune rune
	escaped := false
	runes := []rune(input)

	if len(runes) == 0 {
		return "", nil
	}

	for _, r := range runes {
		if unicode.IsDigit(r) && !escaped {
			if prevRune == 0 {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(prevRune), count))
			prevRune = 0
		} else if r == '\\' && !escaped {
			escaped = true
		} else {
			if prevRune != 0 {
				result.WriteRune(prevRune)
			}
			prevRune = r
			escaped = false
		}
	}

	if escaped {
		return "", ErrWrongEscapeSeq
	}

	if prevRune != 0 {
		result.WriteRune(prevRune)
	}

	return result.String(), nil
}
