package player

import (
	"fmt"
	"go-game-server/function"
	"go-game-server/network"
	"go-game-server/network/protocol/gen/player"
	"google.golang.org/protobuf/proto"
)

type Handler func(*network.SessionPacket)

func (p *Player) AddFriend(packet *network.SessionPacket) {
	req := &player.CSAddFriend{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	if !function.CheckInSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

func (p *Player) DelFriend(packet *network.SessionPacket) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	p.FriendList = function.DelOneInSlice(req.UId, p.FriendList)
}

func (p *Player) ResolveChatMsg(packet *network.SessionPacket) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)
	// TODO 收到消息 转发给客户端
}
