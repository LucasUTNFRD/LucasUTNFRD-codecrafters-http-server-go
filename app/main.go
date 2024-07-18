package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/server"
	"log"
	"net"
)

//TODO refactor code and divide it into different files

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		// fmt.Println("Failed to bind to port 4221")
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.HandleConnection(conn)
	}
}
