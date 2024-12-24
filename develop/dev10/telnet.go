package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func runTelnet() {
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут на подключение")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go-telnet [OPTIONS] HOST PORT")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "connection error: %v\n", err)
		}
		os.Exit(0)
	}()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		_, err := fmt.Fprintln(conn, input.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to write to server: %v\n", err)
			break
		}
	}

	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "failed to read from stdin: %v\n", input.Err())
	}
}
