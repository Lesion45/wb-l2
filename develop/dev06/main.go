package dev06

import "fmt"

func runCut() {
	options := parseArgs()

	lines, err := readInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	for _, line := range lines {
		processedLine := processLine(line, options)
		if processedLine != "" {
			fmt.Println(processedLine)
		}
	}
}

func main() {
	runCut()
}
