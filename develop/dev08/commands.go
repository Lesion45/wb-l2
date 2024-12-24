package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func cd(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("cd: missing argument")
	}
	return os.Chdir(args[0])
}

func pwd() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(cwd)
	return nil
}

func echo(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func kill(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("kill: missing PID")
	}
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "kill: %v\n", err)
	}

	err = proc.Kill()
	if err != nil {
		fmt.Fprintf(os.Stderr, "kill: %v\n", err)
	}
	return nil
}

func ps() error {
	cmd := exec.Command("ps", "-e")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func execute(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
