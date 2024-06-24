package main

import (
	"go-game-server/network"
)

func main() {
	server := network.NewServer("tcp", "127.0.0.1:8023")
	server.Run()
	select {}
}
