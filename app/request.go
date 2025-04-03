package main

import (
	"fmt"
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

		key := endpoint[start+1 : end]
	
		value := strings.TrimPrefix(r.Target, endpoint[:start])
		if valEnd := strings.Index(value, "/"); valEnd != -1 {
			value = value[:valEnd]
		}
		
		r.PathValues[key] = value
	}
}

func NormalizeTarget(r *Request) string {
	target := r.Target
	for key, val := range r.PathValues {
		target = strings.ReplaceAll(target, val, fmt.Sprintf("{%s}", key))
	}

	return target
}
