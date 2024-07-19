package server

import (
	"bufio"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/response"
	"log"
	"net"
	"net/http"
	"strings"
)

func HandleConnection(conn net.Conn, dir string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Println(err)
	}
	responseWriter := response.NewCustomResponseWriter(conn)

	switch {
	case req.URL.Path == "/":
		rootHandler(responseWriter, req)
	case strings.HasPrefix(req.URL.Path, "/files/"):
		fileHandler(responseWriter, req, dir)
	case strings.HasPrefix(req.URL.Path, "/echo/"):
		echoHandler(responseWriter, req)
	case strings.HasPrefix(req.URL.Path, "/user-agent"):
		userAgentHandler(responseWriter, req)
	default:
		http.NotFound(responseWriter, req)
	}
}
