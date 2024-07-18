package request

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

const CRLF = "\r\n"

type Request struct {
	Headers map[string]string
	Body    []byte
	Method  string
	Path    string
}

func ParseRequest(reader *bufio.Reader) (*Request, error) {
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line")
	}

	method := parts[0]
	path := parts[1]

	req := &Request{
		Method:  method,
		Path:    path,
		Headers: make(map[string]string),
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		keyVal := strings.SplitN(line, ":", 2)
		if len(keyVal) != 2 {
			continue
		}
		key := strings.TrimSpace(keyVal[0])
		value := strings.TrimSpace(keyVal[1])
		req.Headers[key] = value
	}

	contentLength := req.Headers["Content-Length"]
	if contentLength != "" {
		length, err := strconv.Atoi(contentLength)
		if err != nil {
			return nil, fmt.Errorf("invalid Content-Length header")
		}
		body := make([]byte, length)
		_, err = io.ReadFull(reader, body)
		if err != nil {
			return nil, err
		}
		req.Body = body
	}
	log.Println(*req)
	return req, nil
}
