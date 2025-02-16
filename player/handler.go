package player

import (
	"fmt"
	"go-game-server/function"
	"go-game-server/network"
	"go-game-server/network/protocol/gen/player"
	"google.golang.org/protobuf/proto"
)

type Handler func(message *network.Message)

func (p *Player) AddFriend(message *network.Message) {
	req := &player.CSAddFriend{}
	err := proto.Unmarshal(message.Data, req)
	if err != nil {
		return
	}
	if !function.CheckInSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

func (p *Player) DelFriend(message *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(message.Data, req)
	if err != nil {
		return
	}
	p.FriendList = function.DelOneInSlice(req.UId, p.FriendList)
}

func (p *Player) ResolveChatMsg(message *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(message.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)
	// TODO 收到消息 转发给客户端
}
