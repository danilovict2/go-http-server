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
	target := NormalizeTarget(req)
	var resp *Response

	if handler, ok := Handlers[target]; ok {
		resp = handler(req)
	} else {
		resp = &Response{
			Protocol:   "HTTP/1.1",
			StatusCode: 404,
			StatusText: "Not Found",
			Headers:    make(map[string]string),
			Body:       "",
		}
	}

	conn.Write(resp.Marshal())
}
