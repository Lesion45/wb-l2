package main

import (
	"os"
	"os/exec"
	"strings"
)

func executeInput(input string) error {
	commands := strings.Split(input, "|")
	if len(commands) == 1 {
		return executeCommand(strings.TrimSpace(commands[0]))
	}
	return executePipeline(commands)
}

func executePipeline(commands []string) error {
	var prevCmd *exec.Cmd

	for i, cmdStr := range commands {
		args := parseCommand(cmdStr)
		cmd := exec.Command(args[0], args[1:]...)

		if i > 0 {
			stdin, err := prevCmd.StdoutPipe()
			if err != nil {
				return err
			}
			cmd.Stdin = stdin
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			return err
		}

		if prevCmd != nil {
			prevCmd.Wait()
		}
		prevCmd = cmd
	}
	return prevCmd.Wait()
}

func executeCommand(input string) error {
	args := parseCommand(input)
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "cd":
		return cd(args[1:])
	case "pwd":
		return pwd()
	case "echo":
		return echo(args[1:])
	case "kill":
		return kill(args[1:])
	case "ps":
		return ps()
	default:
		return execute(args)
	}
}

func parseCommand(input string) []string {
	return strings.Fields(input)
}
