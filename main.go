package main

import "go-game-server/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}
