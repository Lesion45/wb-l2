package main

import (
	"bufio"
	"fmt"
	"os"
)

func runShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("wb-tech$ ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if err := executeInput(input); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
}

func main() {
	runShell()
}
