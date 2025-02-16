package main

import (
	"fmt"
	"go-game-server/network"
	"go-game-server/network/protocol/gen/messageId"
)

type Client struct {
	cli             *network.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[messageId.MessageId]MessageHandler
	console         *ClientConsole
	chInput         chan *InputParam
}

func NewClient() *Client {
	c := &Client{
		cli:             network.NewClient(":8023"),
		inputHandlers:   make(map[string]InputHandler),
		messageHandlers: make(map[messageId.MessageId]MessageHandler),
		console:         NewClientConsole(),
	}
	c.cli.OnMessage = c.OnMessage
	c.cli.ChMsg = make(chan *network.Message, 1)
	c.chInput = make(chan *InputParam, 1)
	c.console.chInput = c.chInput
	return c
}

func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput:
				fmt.Printf("cmd:%s,param:%v	<<<\t \n", input.Command, input.Param)
				inputHandler := c.inputHandlers[input.Command]
				if inputHandler != nil {
					inputHandler(input)
				}
			}
		}
	}()
	go c.console.Run()
	go c.cli.Run()
}

func (c *Client) OnMessage(packet *network.ClientPacket) {
	if handler, ok := c.messageHandlers[messageId.MessageId(packet.Msg.ID)]; ok {
		handler(packet)
	}
}
