package main

import (
	"flag"
	"github.com/codecrafters-io/http-server-starter-go/app/internal/server"
	"log"
	"net"
	"os"
)

// TODO implement /file/{filename} endpoint
func main() {
	dirPtr := flag.String("directory", "/tmp/", "specify directory for files")
	flag.Parse()

	if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
		log.Fatalf("directory does not exist: %s", *dirPtr)
	}

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
		go server.HandleConnection(conn, *dirPtr)
	}
}
