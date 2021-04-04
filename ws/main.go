package main

import (
	"log"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	log.Println("Started at 8080")
	if err != nil {
		// handle error
	}

	// Prepare handshake header writer from http.Header mapping.

	u := ws.Upgrader{}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		_, err = u.Upgrade(conn)
		if err != nil {
			log.Printf("upgrade error: %s", err)
		}
		go func() {
			defer conn.Close()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					// handle error
				}
			}
		}()
	}
}
