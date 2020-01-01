package cg

import "fmt"

type Player struct {
	Name         string
	Level        int
	Experience   int
	Room         int
	MessageQueue chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}
	go func(p *Player) {
		for {
			message := <-p.MessageQueue
			fmt.Println(p.Name, "received message: ", message.Content)
		}
	}(player)
	return player
}
