package main

import (
	"go-game-server/network"
)

func main() {
	server := network.NewServer("tcp6", ":8023")
	server.Run()
	select {
		
	}
}