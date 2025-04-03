package main

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	l net.Listener
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Server is listening on port 4221...")

	server := &Server{
		l: l,
	}
	server.Accept()
}

func (s *Server) Accept() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()

	rawRequest := make([]byte, 1024)
	_, err := conn.Read(rawRequest)
	if err != nil {
		fmt.Println("Error reading the request: ", err.Error())
		os.Exit(1)
	}

	req := Unmarshal(rawRequest)
	if req.Target != "/" {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	}
}
