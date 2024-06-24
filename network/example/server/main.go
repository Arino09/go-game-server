package main

import (
	"go-game-server/network"
)

func main() {
	server := network.NewServer("tcp", ":8023")
	server.Run()
	select {}
}
