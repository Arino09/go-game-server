package main

import (
	"go-game-server/network"
)

func main() {
	server := network.NewServer(":8023")
	server.Run()
	select {}
}
