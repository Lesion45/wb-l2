package main

import (
	"flag"
	"fmt"
	"os"
)

func sorterRun() {
	options := parseFlags()
	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "Specify file name")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	lines, err := readLines(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read file:", err)
		os.Exit(1)
	}

	sorter := NewSorter(lines, options)
	if options.CheckSorted {
		if sorter.IsSorted() {
			fmt.Println("file is already sorted")
		} else {
			fmt.Println("file isn't sorted")
		}
		return
	}

	sortedLines := sorter.Sort()
	for _, line := range sortedLines {
		fmt.Println(line)
	}
}

func main() {
	sorterRun()
}
