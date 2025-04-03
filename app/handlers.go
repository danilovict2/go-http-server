package main

import (
	"strconv"
)

type Handler func(req *Request) *Response

var Handlers = map[string]Handler{
	"/":           root,
	"/echo/{str}": echo,
	"/user-agent": userAgent,
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
