package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1)
	for {
		_, err := io.ReadFull(conn, buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		_, err = conn.Write(buf)
		if err != nil {
			log.Fatal(err)
		}
	}
}
