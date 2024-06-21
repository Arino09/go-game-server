package manager

import (
	"go-game-server/player"
)

// 玩家管理器
type PlayerMgr struct {
	playerList map[uint64]player.Player
	addPCh chan player.Player
}

func (pm *PlayerMgr) Add(p player.Player) {
	pm.playerList[p.UId] = p
	go p.Run()
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <- pm.addPCh:
			pm.Add(p)
		}
	}
}
