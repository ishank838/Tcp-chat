package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Printf("Started Listening on :8888")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("unable to accept new Connection")
			continue
		}

		go s.newClient(conn)
	}
}
