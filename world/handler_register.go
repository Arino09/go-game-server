package world

import "go-game-server/network/protocol/gen/messageId"

func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[messageId.MessageId_CSCreateUser] = mm.CreatePlayer
	mm.Handlers[messageId.MessageId_CSLogin] = mm.UserLogin
}
