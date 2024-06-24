package player

import (
	"fmt"
	"go-game-server/chat"
	"go-game-server/function"
)

type Handler func(interface{})

func (p *Player) AddFriend(data interface{}) {
	fId := data.(uint64)
	if !function.CheckInSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

func (p *Player) DelFriend(data interface{}) {
	fId := data.(uint64)
	p.FriendList = function.DelOneInSlice(fId, p.FriendList)
}

func (p *Player) ResolveChatMsg(data interface{}) {
	chatMsg := data.(chat.Msg)
	fmt.Println(chatMsg)
	// TODO 收到消息 转发给客户端
}
