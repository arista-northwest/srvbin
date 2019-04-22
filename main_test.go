package main

import (
	"bytes"
	"log"
	"net"
	"testing"
)

var (
	server Server
	addr   string = ":50002"
)

func init() {
	server, err := NewServer(addr)
	if err != nil {
		log.Println("error starting TCP server")
		return
	}

	go func() {
		server.Run()
	}()
}

func TestTCPServer_Running(t *testing.T) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error("Failed to connect to server: ", err)
	}

	defer conn.Close()
}

func TestTCPServer_Echo(t *testing.T) {
	tests := []struct {
		descr   string
		payload []byte
	}{
		{"Send a string...", []byte("ECHO hello\n")},
		{"Send another string...", []byte("ECHO hello?\n")},
	}

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	for _, tc := range tests {
		t.Run(tc.descr, func(t *testing.T) {
			if _, err := conn.Write(tc.payload); err != nil {
				t.Error("could not write payload to server:", err)
			}

			out := make([]byte, 1024)

			_, err := conn.Read(out)

			if err != nil {
				t.Error("could not read from connection")
			}

			if bytes.Compare(out, tc.payload) == 0 {
				t.Error("response did match expected output")
			}

		})
	}

	if err != nil {
		t.Error("Failed to connect: ", err)
	}
}
