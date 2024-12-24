package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func runGrep() {
	after := flag.Int("A", 0, "Печатать +N строк после совпадения")
	before := flag.Int("B", 0, "Печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "Количество строк")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	invertMatch := flag.Bool("v", false, "Вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "Точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "Напечатать номер строки")

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		os.Exit(1)
	}

	pattern := args[0]
	filename := ""
	if len(args) > 1 {
		filename = args[1]
	}

	lines, err := ReadInput(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	options := FilterOptions{
		After:       *after,
		Before:      *before,
		Context:     *context,
		Count:       *count,
		IgnoreCase:  *ignoreCase,
		InvertMatch: *invertMatch,
		Fixed:       *fixed,
		LineNum:     *lineNum,
		Pattern:     pattern,
	}

	results := Filter(lines, options)
	PrintResults(results, options)
}

func PrintResults(results []FilterResult, options FilterOptions) {
	for _, result := range results {
		if options.LineNum {
			fmt.Printf("%d:", result.LineNum)
		}
		fmt.Println(result.Line)
	}
}

func ReadInput(filename string) ([]string, error) {
	var lines []string
	var scanner *bufio.Scanner

	if filename == "" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
