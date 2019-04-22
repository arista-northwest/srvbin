package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
)

type Server interface {
	Run() error
	Close() error
}

type TCPServer struct {
	addr   string
	server net.Listener
}

func NewServer(addr string) (Server, error) {
	return &TCPServer{
		addr: addr,
	}, nil
}

func (t *TCPServer) Run() (err error) {
	t.server, err = net.Listen("tcp", t.addr)
	if err != nil {
		return
	}
	defer t.Close()

	for {
		conn, err := t.server.Accept()
		fmt.Print("Got a new connection\n")
		if err != nil {
			err = errors.New("could not accept connection")
			break
		}

		if conn == nil {
			err = errors.New("could not create connection")
			break
		}
		go t.handleConnection(conn)
	}
	return
}

func (t *TCPServer) Close() (err error) {
	return t.server.Close()
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {

		req, err := rw.ReadString('\n')

		if err != nil {
			rw.WriteString("failed to read input")
			rw.Flush()
			return
		}

		fmt.Printf("Go new data: %s", req)

		for {
			_, err = rw.WriteString(fmt.Sprintf("%s", req))
			if err != nil {
				fmt.Print("Client went away.\n")
				break
			}
			rw.Flush()
		}

	}
}

func main() {
	server, err := NewServer(":1123")
	if err != nil {
		log.Println("error starting TCP server")
		return
	}

	fmt.Print("Starting server...\n")
	server.Run()
}
