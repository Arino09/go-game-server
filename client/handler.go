package main

import (
	"fmt"
	"go-game-server/network"
)

type MessageHandler func(packet *network.ClientPacket)

type InputHandler func(param *InputParam)

func (c *Client) CreatePlayer(param *InputParam) {
	//id := c.GetMessageIdByCmd(param.Command)
	//if len(param.Param) != 2 {
	//    return
	//}
	//msg := &player.CSCreateUser{
	//    UserName: param.Param[0],
	//    Password: param.Param[1],
	//}
	//c.cli.ChMsg <- &network.Message{}
}

func (c *Client) OnCreatePlayerRsp(packet *network.ClientPacket) {

}

func (c *Client) Login(param *InputParam) {
	fmt.Printf("Login input handler print: ")
	fmt.Println(param.Command)
	fmt.Println(param.Param)
}

func (c *Client) OnLoginRsp(packet *network.ClientPacket) {

}
func (c *Client) AddFriend(param *InputParam) {

}

func (c *Client) OnAddFriendRsp(packet *network.ClientPacket) {

}

func (c *Client) DelFriend(param *InputParam) {

}

func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {

}

func (c *Client) SendChatMsg(param *InputParam) {

}

func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {

}
