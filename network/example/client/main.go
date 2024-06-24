package main

import (
	"go-game-server/network"
)

func main() {
	client := network.NewClient("127.0.0.1:8023")
	client.Run()
	select {}
}
