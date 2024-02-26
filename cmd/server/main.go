package main

import (
	"fmt"
	"memtable/internal"

	tcp "github.com/rayaan01/tcp-server"
)

func main() {
	server, err := tcp.CreateServer("localhost", 8080)
	if err != nil {
		fmt.Println("Server could not start", err)
		return
	}
	server.AcceptConnections(internal.Handler)
}
