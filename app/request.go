package main

import (
	"strings"
)

type Request struct {
	Method     string
	Target     string
	Protocol   string
	PathValues map[string]string
}

func Unmarshal(rawRequest []byte) *Request {
	parts := strings.Split(string(rawRequest), "\r\n")
	line := strings.Split(parts[0], " ")
	req := &Request{
		Method:     line[0],
		Target:     line[1],
		Protocol:   line[2],
		PathValues: make(map[string]string),
	}

	req.setPathValues()

	return req
}

func (r *Request) setPathValues() {
	for endpoint := range Handlers {
		start := strings.Index(endpoint, "{")
		if start == -1 {
			continue
		}

		end := strings.Index(endpoint[start:], "}")
		if end == -1 {
			continue
		}		
		end += start

		if !strings.HasPrefix(r.Target, endpoint[:start]) || !strings.HasSuffix(r.Target, endpoint[end+1:]) {
			continue
		}

		name := endpoint[start+1:end]
		value := strings.TrimSuffix(strings.TrimPrefix(r.Target, endpoint[:start]), endpoint[end+1:])
		r.PathValues[name] = value
	}
}
