package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Response struct {
	Protocol   string
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       string
}

func (r *Response) Marshal() []byte {
	ret := make([]byte, 0)
	ret = append(ret, r.marshalStatusLine()...)
	ret = append(ret, '\r', '\n')

	ret = append(ret, r.marshalHeaders()...)
	ret = append(ret, '\r', '\n')

	ret = append(ret, []byte(r.Body)...)

	return ret
}

func (r *Response) marshalStatusLine() []byte {
	ret := make([]byte, 0)
	ret = append(ret, []byte(r.Protocol)...)
	ret = append(ret, ' ')
	ret = append(ret, []byte(strconv.Itoa(r.StatusCode))...)
	ret = append(ret, ' ')
	ret = append(ret, []byte(r.StatusText)...)

	return ret
}

func (r *Response) marshalHeaders() []byte {
	ret := make([]byte, 0)
	for key, value := range r.Headers {
		header := fmt.Sprintf("%s: %s", key, value)
		ret = append(ret, []byte(header)...)
		ret = append(ret, '\r', '\n')
	}

	return ret
}

func (resp *Response) TryCompress(req *Request) {
	if val, ok := req.Headers["accept-encoding"]; ok && slices.Contains(strings.Split(val, ", "), "gzip") {
		var b bytes.Buffer
		w := gzip.NewWriter(&b)

		if _, err := w.Write([]byte(resp.Body)); err != nil {
			fmt.Println("Failed to compress response body: ", err.Error())
			return
		}
		w.Close()

		resp.Body = b.String()
		resp.Headers["Content-Encoding"] = "gzip"
		resp.Headers["Content-Length"] = strconv.Itoa(len(resp.Body))
	}
}
