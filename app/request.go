package main

import "strings"

type Request struct {
	Method  string
	Target  string
	Version string
}

func Unmarshal(rawRequest []byte) *Request {
	parts := strings.Split(string(rawRequest), "\r\n")
	line := strings.Split(parts[0], " ")

	return &Request{
		Method:  line[0],
		Target:  line[1],
		Version: line[2],
	}
}
