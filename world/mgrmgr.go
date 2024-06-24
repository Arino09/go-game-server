package world

import (
	"go-game-server/manager"
)

type MgrMgr struct {
	Pm manager.PlayerMgr
}

var MM *MgrMgr

func NewMgrMgr() *MgrMgr {
	return &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
}
