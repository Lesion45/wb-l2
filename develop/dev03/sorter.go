package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Sorter struct {
	lines   []string
	options SortOptions
}

type SortOptions struct {
	Column         int
	Numeric        bool
	Reverse        bool
	Unique         bool
	Month          bool
	IgnoreTrailing bool
	CheckSorted    bool
	HumanReadable  bool
}

func parseFlags() SortOptions {
	column := flag.Int("k", 0, "Номер колонки для сортировки (1-based)")
	numeric := flag.Bool("n", false, "Числовая сортировка")
	reverse := flag.Bool("r", false, "Обратный порядок сортировки")
	unique := flag.Bool("u", false, "Удаление дубликатов")
	month := flag.Bool("M", false, "Сортировка по месяцам")
	ignoreTrailing := flag.Bool("b", false, "Игнорировать хвостовые пробелы")
	checkSorted := flag.Bool("c", false, "Проверка на упорядоченность")
	humanReadable := flag.Bool("h", false, "Сортировка с учетом суффиксов")

	flag.Parse()

	return SortOptions{
		Column:         *column,
		Numeric:        *numeric,
		Reverse:        *reverse,
		Unique:         *unique,
		Month:          *month,
		IgnoreTrailing: *ignoreTrailing,
		CheckSorted:    *checkSorted,
		HumanReadable:  *humanReadable,
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func NewSorter(lines []string, options SortOptions) *Sorter {
	if options.IgnoreTrailing {
		trimTrailingSpaces(lines)
	}
	return &Sorter{
		lines:   lines,
		options: options,
	}
}

func (s *Sorter) Sort() []string {
	if s.options.Month {
		monthOrder := map[string]int{
			"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6,
			"Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
		}
		sort.Slice(s.lines, func(i, j int) bool {
			return compareMonths(getColumn(s.lines[i], s.options.Column), getColumn(s.lines[j], s.options.Column), monthOrder, s.options.Reverse)
		})
	} else if s.options.Numeric || s.options.HumanReadable {
		sort.Slice(s.lines, func(i, j int) bool {
			return compareNumeric(getColumn(s.lines[i], s.options.Column), getColumn(s.lines[j], s.options.Column), s.options.HumanReadable, s.options.Reverse)
		})
	} else {
		sort.Slice(s.lines, func(i, j int) bool {
			return compareStrings(getColumn(s.lines[i], s.options.Column), getColumn(s.lines[j], s.options.Column), s.options.Reverse)
		})
	}

	if s.options.Unique {
		s.lines = removeDuplicates(s.lines)
	}
	return s.lines
}

func (s *Sorter) IsSorted() bool {
	for i := 0; i < len(s.lines)-1; i++ {
		if compareLines(s.lines[i], s.lines[i+1], s.options.Column, s.options.Numeric, s.options.Reverse, s.options.Month, s.options.HumanReadable) > 0 {
			return false
		}
	}
	return true
}

func trimTrailingSpaces(lines []string) {
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " ")
	}
}

func getColumn(line string, k int) string {
	if k <= 0 {
		return line
	}
	fields := strings.Fields(line)
	if k > len(fields) {
		return ""
	}
	return fields[k-1]
}

func compareStrings(a, b string, reverse bool) bool {
	if reverse {
		return a > b
	}
	return a < b
}

func compareNumeric(a, b string, humanReadable, reverse bool) bool {
	aNum, _ := parseNumeric(a, humanReadable)
	bNum, _ := parseNumeric(b, humanReadable)
	if reverse {
		return aNum > bNum
	}
	return aNum < bNum
}

func compareMonths(a, b string, monthOrder map[string]int, reverse bool) bool {
	aVal, aOk := monthOrder[a]
	bVal, bOk := monthOrder[b]
	if !aOk || !bOk {
		return false
	}
	if reverse {
		return aVal > bVal
	}
	return aVal < bVal
}

func compareLines(a, b string, k int, n, r, M, h bool) int {
	aCol := getColumn(a, k)
	bCol := getColumn(b, k)

	if M {
		monthOrder := map[string]int{
			"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6,
			"Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
		}
		return compareMonthOrder(aCol, bCol, monthOrder, r)
	} else if n || h {
		return compareNumericOrder(aCol, bCol, h, r)
	}
	return compareStringOrder(aCol, bCol, r)
}

func compareMonthOrder(a, b string, monthOrder map[string]int, reverse bool) int {
	aVal, aOk := monthOrder[a]
	bVal, bOk := monthOrder[b]
	if !aOk || !bOk {
		return strings.Compare(a, b)
	}
	if reverse {
		return bVal - aVal
	}
	return aVal - bVal
}

func compareNumericOrder(a, b string, humanReadable, reverse bool) int {
	aNum, _ := parseNumeric(a, humanReadable)
	bNum, _ := parseNumeric(b, humanReadable)
	if reverse {
		if aNum > bNum {
			return -1
		} else if aNum < bNum {
			return 1
		}
	} else {
		if aNum < bNum {
			return -1
		} else if aNum > bNum {
			return 1
		}
	}
	return 0
}

func compareStringOrder(a, b string, reverse bool) int {
	if reverse {
		return strings.Compare(b, a)
	}
	return strings.Compare(a, b)
}

func parseNumeric(s string, humanReadable bool) (float64, error) {
	if humanReadable {
		return parseHumanReadableNumber(s)
	}
	return strconv.ParseFloat(s, 64)
}

func parseHumanReadableNumber(s string) (float64, error) {
	multiplier := 1.0
	if len(s) > 0 {
		switch s[len(s)-1] {
		case 'K', 'k':
			multiplier = 1e3
		case 'M', 'm':
			multiplier = 1e6
		case 'G', 'g':
			multiplier = 1e9
		}
		s = s[:len(s)-1]
	}
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return num * multiplier, nil
}

func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	var uniqueLines []string
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}
