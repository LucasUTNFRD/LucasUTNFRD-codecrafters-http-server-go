package response

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

const (
	CRLF    = "\r\n"
	VERSION = "HTTP/1.1"
)

type CustomResponseWriter struct {
	conn   net.Conn
	header http.Header
	status int
	buf    *bufio.Writer
}

func NewCustomResponseWriter(conn net.Conn) *CustomResponseWriter {
	return &CustomResponseWriter{
		conn:   conn,
		buf:    bufio.NewWriter(conn),
		header: make(http.Header),
	}
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.header
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func (w *CustomResponseWriter) Write(data []byte) (int, error) {
	// Write status line
	if w.status == 0 {
		w.status = http.StatusOK
	}
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", w.status, http.StatusText(w.status))
	_, err := w.buf.WriteString(statusLine)
	if err != nil {
		return 0, err
	}

	// Write headers
	for key, values := range w.header {
		for _, value := range values {
			_, err = w.buf.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
			if err != nil {
				return 0, err
			}
		}
	}
	_, err = w.buf.WriteString("\r\n")
	if err != nil {
		return 0, err
	}

	// Write body
	n, err := w.buf.Write(data)
	if err != nil {
		return n, err
	}
	err = w.buf.Flush()
	return n, err
}
