package websocket

import (
	"fmt"
)

//Room - data from room
type Room struct {
	RoomID  string
	Clients map[*Client]bool
	Host    *Client
	Execute chan Command
}

//Command - the formula to issue commands to a room.
type Command struct {
	Sender  *Client
	Command string
	Type    int
}

//NewRoom - constructor for new rooms
func NewRoom(id string, host *Client) *Room {

	clientMap := make(map[*Client]bool)
	clientMap[host] = true

	return &Room{
		RoomID:  id,
		Clients: clientMap,
		Host:    host,
		Execute: make(chan Command),
	}
}

//May or may not to be exported, need to check that out.
func (room *Room) start() {
	for {
		select {
		case command := <-room.Execute:
			fmt.Println(command)
		}
	}
}
