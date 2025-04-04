package main

import (
	"os"
	"strconv"
)

type Handler func(req *Request) *Response

var Handlers = map[string]Handler{
	"GET /":                  root,
	"GET /echo/{str}":        echo,
	"GET /user-agent":        userAgent,
	"GET /files/{filename}":  getFile,
	"POST /files/{filename}": createFile,
}

func root(req *Request) *Response {
	return &Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers:    make(map[string]string),
		Body:       "",
	}
}

func echo(req *Request) *Response {
	headers := make(map[string]string)
	headers["Content-Type"] = "text/plain"
	headers["Content-Length"] = strconv.Itoa(len(req.PathValues["str"]))

	return &Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers:    headers,
		Body:       req.PathValues["str"],
	}
}

func userAgent(req *Request) *Response {
	headers := make(map[string]string)
	headers["Content-Type"] = "text/plain"
	headers["Content-Length"] = strconv.Itoa(len(req.Headers["user-agent"]))

	return &Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers:    headers,
		Body:       req.Headers["user-agent"],
	}
}

func getFile(req *Request) *Response {
	content, err := os.ReadFile(directory + req.PathValues["filename"])
	if os.IsNotExist(err) {
		return &Response{
			Protocol:   "HTTP/1.1",
			StatusCode: 404,
			StatusText: "Not Found",
			Headers:    make(map[string]string),
			Body:       "",
		}
	} else if err != nil {
		return &Response{
			Protocol:   "HTTP/1.1",
			StatusCode: 500,
			StatusText: "Internal Server Error",
			Headers:    make(map[string]string),
			Body:       "",
		}
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	headers["Content-Length"] = strconv.Itoa(len(content))

	return &Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers:    headers,
		Body:       string(content),
	}
}

func createFile(req *Request) *Response {
	f, err := os.Create(directory + req.PathValues["filename"])
	if err != nil {
		return &Response{
			Protocol:   "HTTP/1.1",
			StatusCode: 500,
			StatusText: "Internal Server Error",
			Headers:    make(map[string]string),
			Body:       "",
		}
	}
	defer f.Close()

	if _, err := f.WriteString(req.Body); err != nil {
		return &Response{
			Protocol:   "HTTP/1.1",
			StatusCode: 500,
			StatusText: "Internal Server Error",
			Headers:    make(map[string]string),
			Body:       "",
		}
	}

	return &Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 201,
		StatusText: "Created",
		Headers:    make(map[string]string),
		Body:       "",
	}
}
