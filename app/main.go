package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Server struct {
	l net.Listener
}

var directory string

func main() {
	flag.Parse()

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

func init() {
	flag.StringVar(&directory, "directory", "/tmp/", "Specifies the directory where the files are stored, as an absolute path")
}

func (s *Server) Accept() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()

	for {
		rawRequest := make([]byte, 1024)
		n, err := conn.Read(rawRequest)
		if errors.Is(err, io.EOF) {
			fmt.Println("Client closed the connections:", conn.RemoteAddr())
			break
		} else if err != nil {
			fmt.Println("Error reading the request: ", err.Error())
			break
		}

		req := Unmarshal(rawRequest[:n])
		target := fmt.Sprintf("%s %s", req.Method, NormalizeTarget(req))
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

		resp.TryCompress(req)
		conn.Write(resp.Marshal())

		if val, ok := req.Headers["connection"]; req.Protocol != "HTTP/1.1" || (ok && val == "close") {
			break	
		}
	}
}
