package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClientConsole struct {
	chInput chan *InputParam
}

type InputParam struct {
	Command string
	Param   []string
}

func NewClientConsole() *ClientConsole {
	c := &ClientConsole{}
	return c
}

func (c *ClientConsole) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error, check your input and try again!")
			continue
		}
		replace := strings.Replace(readString, "\n", "", -1)
		split := strings.Split(replace, " ")
		if len(split) == 0 {
			fmt.Println("Input error, check your input and try again!")
			continue
		}
		in := &InputParam{
			Command: split[0],
			Param:   split[1:],
		}
		c.chInput <- in
	}
}
