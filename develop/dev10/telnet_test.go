package main

import (
	"net"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Logf("Mock server error: %v", err)
			return
		}
		defer conn.Close()

		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			conn.Write(buf[:n])
		}
	}()

	go func() {
		err := RunTelnetClient("127.0.0.1", "8080", 5*time.Second)
		if err != nil {
			t.Errorf("Telnet client error: %v", err)
		}
	}()

	time.Sleep(2 * time.Second)
}

func RunTelnetClient(host, port string, timeout time.Duration) error {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte("test\n"))
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return err
	}
	return nil
}
