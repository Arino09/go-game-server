package world

import (
	"go-game-server/network"
	"go-game-server/player"
)

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	newPlayer := player.NewPlayer()
	newPlayer.UId = 111
	newPlayer.HandlerParamCh = message.Sess.WriteCh
	message.Sess.IsPlayerOnline = true
	newPlayer.Run()
}
