package server

import (
	"bufio"
	"log"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/response"
)

func HandleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	reader := bufio.NewReader(conn)
	req, err := request.ParseRequest(reader)
	if err != nil {
		log.Println("Error parsing request:", err)
		return
	}

	res := response.NewResponse(req)
	log.Println("response", *res)
	err = res.SendResponse(conn)
	if err != nil {
		log.Println("Error sending response:", err)
	}
}
