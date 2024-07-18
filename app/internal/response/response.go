package response

import (
	"fmt"
	"net"
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/request"
)

const CRLF = "\r\n"

const (
	OK       = 200
	NotFound = 404
)

type Response struct {
	Status  int
	Message string
	Body    []byte
	Headers map[string]string
}

func NewResponse(req *request.Request) *Response {
	var status int
	var message string
	var body string

	if req.Path == "/" {
		status = OK
		message = "OK"
	} else {
		status = NotFound
		message = "Not Found"
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "text/plain"
	headers["Content-Length"] = strconv.Itoa(len(body))

	return &Response{
		Status:  status,
		Message: message,
		Body:    []byte(body),
		Headers: headers,
	}
}

func (res *Response) SendResponse(conn net.Conn) error {
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s%s", res.Status, res.Message, CRLF)
	headers := ""
	for k, v := range res.Headers {
		headers += fmt.Sprintf("%s: %s%s", k, v, CRLF)
	}

	response := statusLine + headers + CRLF + string(res.Body)
	_, err := conn.Write([]byte(response))
	return err
}
